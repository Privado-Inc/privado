package docker

import (
	"bufio"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/Privado-Inc/privado/pkg/config"
	"github.com/Privado-Inc/privado/pkg/utils"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/jsonmessage"
	"github.com/moby/term"

	// "github.com/docker/docker/ne"
	"github.com/docker/go-connections/nat"
)

func getDefaultDockerClient() (*client.Client, error) {
	client, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}

	return client, nil
}

func getBaseContainerConfig(image string) *container.Config {
	config := &container.Config{
		Image:        image,
		AttachStdin:  true,
		AttachStdout: true,
		AttachStderr: true,
		OpenStdin:    true,
		Tty:          true,
		ExposedPorts: nat.PortSet{
			nat.Port(config.AppConfig.Container.WebPort): {},
		},
	}
	return config
}

func getContainerHostConfig(volumes containerVolumes, ports containerPorts) *container.HostConfig {
	hostConfig := &container.HostConfig{}

	if volumes.licenseVolumeEnabled {
		hostConfig.Mounts = append(
			hostConfig.Mounts,
			mount.Mount{
				Type:     "bind",
				Source:   volumes.licenseVolumeHost,
				Target:   config.AppConfig.Container.LicenseVolumeDir,
				ReadOnly: true,
			},
		)
	}
	if volumes.sourceCodeVolumeEnabled {
		hostConfig.Mounts = append(
			hostConfig.Mounts,
			mount.Mount{
				Type:   "bind",
				Source: volumes.sourceCodeVolumeHost,
				Target: config.AppConfig.Container.SourceCodeVolumeDir,
			},
		)
	}

	if ports.webPortEnabled {
		hostConfig.PortBindings = map[nat.Port][]nat.PortBinding{
			nat.Port(config.AppConfig.Container.WebPort): {
				{
					HostIP:   "0.0.0.0",
					HostPort: fmt.Sprint(ports.webPortHost),
				},
			},
		}
	}

	return hostConfig
}

func PullLatestImage(image string, client *client.Client) (err error) {
	if client == nil {
		client, err = getDefaultDockerClient()
		if err != nil {
			return err
		}
	}

	auth := types.AuthConfig{
		Username:      "AWS",
		Password:      os.Getenv("PRIVADO_ECR_ACCESS_CODE"),
		ServerAddress: "638117407428.dkr.ecr.region.amazonaws.com",
	}
	authBytes, _ := json.Marshal(auth)
	authBase64 := base64.URLEncoding.EncodeToString(authBytes)

	ctx := context.Background()

	fmt.Println("\n> Pulling the latest image:", image)
	reader, err := client.ImagePull(ctx, image, types.ImagePullOptions{RegistryAuth: authBase64})
	if err != nil {
		return err
	}

	id, isTerm := term.GetFdInfo(os.Stdout)
	_ = jsonmessage.DisplayJSONMessagesStream(reader, os.Stdout, id, isTerm, nil)

	defer reader.Close()
	io.Copy(os.Stdout, reader)

	return nil
}

//lint:ignore U1000 Ignore unused function warning
func attachContainerOutput(client *client.Client, ctx context.Context, containerId string) (*bufio.Reader, error) {
	waiter, err := client.ContainerAttach(ctx, containerId, types.ContainerAttachOptions{
		Stderr: true,
		Stdout: true,
		Stdin:  true,
		Stream: true,
		Logs:   true,
	})

	if err != nil {
		return nil, err
	}

	// attach io by default for now
	// unsure of the impact yet
	go io.Copy(waiter.Conn, os.Stdin)

	return waiter.Reader, err
}

func processAttachedContainerOutput(reader *bufio.Reader, attachStdOut bool, outputMatchList []string, matchFn func(string)) {
	if attachStdOut {
		go io.Copy(os.Stdout, reader)
		go io.Copy(os.Stderr, reader)
	}

	if len(outputMatchList) > 0 {
		go func() {
			line, _ := reader.ReadString('\n')
			for _, matchStr := range outputMatchList {
				if strings.Contains(line, matchStr) {
					matchFn(strings.TrimSuffix(line, "\n"))
					return
				}
			}
		}()
	}
}

func WaitForContainer(client *client.Client, ctx context.Context, containerId string) error {
	statusCh, errCh := client.ContainerWait(ctx, containerId, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			return err
		}
	case <-statusCh:
	}

	return nil
}

func RemoveContainerForcefully(client *client.Client, ctx context.Context, containerId string) error {
	return client.ContainerRemove(
		ctx,
		containerId,
		types.ContainerRemoveOptions{
			RemoveVolumes: true,
			Force:         true,
		},
	)
}

func StopContainer(client *client.Client, ctx context.Context, containerId string) error {
	return client.ContainerStop(ctx, containerId, nil)
}

func RunImage(opts ...RunImageOption) error {
	runOptions := newRunImageHandler(opts)

	ctx := context.Background()

	client, err := getDefaultDockerClient()
	if err != nil {
		return err
	}

	// Pull image
	image := config.AppConfig.Container.ImageURL
	if err := PullLatestImage(image, client); err != nil {
		return err
	}

	// Generate container configurations
	containerConfig := getBaseContainerConfig(image)
	containerConfig.Cmd = runOptions.args
	hostConfig := getContainerHostConfig(runOptions.volumes, runOptions.ports)

	// Create container
	creationResponse, err := client.ContainerCreate(ctx, containerConfig, hostConfig, nil, nil, "")
	if err != nil {
		return err
	}
	if len(creationResponse.Warnings) > 0 {
		fmt.Println("\n> Encountered warnings:")
		for i, warn := range creationResponse.Warnings {
			fmt.Println(i+1, warn)
		}
	}

	// always remove the container in the end
	defer RemoveContainerForcefully(client, ctx, creationResponse.ID)

	quitProgressBarChannel := make(chan bool, 1)

	// Attach input/output streams with container
	if runOptions.attachOutput || len(runOptions.exitErrorMessages) > 0 {
		// processContainerOutput(attachStdIO, runOnMatch)
		reader, err := attachContainerOutput(client, ctx, creationResponse.ID)
		if err != nil {
			return err
		}

		processAttachedContainerOutput(reader, runOptions.attachOutput, runOptions.exitErrorMessages, func(err string) {
			// Error on output
			quitProgressBarChannel <- true
			fmt.Println("\n> Some error occurred")
			if err != "" {
				// reset any color from internal process
				fmt.Println("Find more details below:\n", err, "\033[0m")
			}
			fmt.Println("\n> Please try again or open an issue here: ", config.AppConfig.PrivadoRepository)
			fmt.Println("\n> Terminating..")
			RemoveContainerForcefully(client, ctx, creationResponse.ID)
		})
	}

	// Start container
	fmt.Println("\n> Starting container with the latest image")
	fmt.Println("> Container ID:", creationResponse.ID)
	if err := client.ContainerStart(ctx, creationResponse.ID, types.ContainerStartOptions{}); err != nil {
		return err
	}

	// Setup interrupt fns if enabled
	if runOptions.setupInterrupt {
		// Listen for interrupt, clear signal after execution
		// Remove container when received
		// All cleanup here: The process ends after this
		// and defer functions are not executed
		sgn := utils.RunOnCtrlC(func() {
			quitProgressBarChannel <- true
			fmt.Println("\n> Received interrupt signal")
			RemoveContainerForcefully(client, ctx, creationResponse.ID)
		})
		defer utils.ClearSignals(sgn)
	}

	if runOptions.ports.webPortEnabled && runOptions.spawnWebBrowser {
		completeProgressChannel := make(chan bool)
		if runOptions.progressLoader {
			go utils.RenderProgressSpinnerWithMessages(
				completeProgressChannel,
				quitProgressBarChannel,
				runOptions.duringLoadMessages,
				runOptions.afterLoadMessages,
			)
		}
		go utils.WaitAndOpenURL(fmt.Sprintf("http://localhost:%d", runOptions.ports.webPortHost), completeProgressChannel, 6)
	}

	// Image output after this point
	fmt.Println("\n> Waiting for process to complete:")

	// wait for container to stop (automatically or by interrupt)
	if err := WaitForContainer(client, ctx, creationResponse.ID); err != nil {
		return err
	}

	return nil
}
