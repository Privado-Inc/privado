package docker

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"os"

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

type ContainerVolumes struct {
	LicenseVolumeEnabled, SourceCodeVolumeEnabled bool
	LicenseVolumeHost, SourceCodeVolumeHost       string
}

type ContainerPorts struct {
	WebPortEnabled bool
	WebPortHost    int
}

type RunImageOptions struct {
	Args                            []string
	Volumes                         *ContainerVolumes
	Ports                           *ContainerPorts
	SetupInterrupt, SpawnWebBrowser bool
}

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
			nat.Port("80/tcp"): {},
		},
	}
	return config
}

func getContainerHostConfig(volumes *ContainerVolumes, ports *ContainerPorts) *container.HostConfig {
	hostConfig := &container.HostConfig{}

	if volumes != nil && volumes.LicenseVolumeEnabled {
		hostConfig.Mounts = append(
			hostConfig.Mounts,
			mount.Mount{
				Type:     "bind",
				Source:   volumes.LicenseVolumeHost,
				Target:   "/tmp/license.json",
				ReadOnly: true,
			},
		)
	}
	if volumes != nil && volumes.SourceCodeVolumeEnabled {
		hostConfig.Mounts = append(
			hostConfig.Mounts,
			mount.Mount{
				Type:   "bind",
				Source: volumes.SourceCodeVolumeHost,
				Target: "/app/code",
			},
		)
	}

	if ports != nil && ports.WebPortEnabled {
		hostConfig.PortBindings = map[nat.Port][]nat.PortBinding{
			nat.Port("80/tcp"): {
				{
					HostIP:   "0.0.0.0",
					HostPort: fmt.Sprint(ports.WebPortHost),
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
		Password:      "=",
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
func attachStdioWithContainer(client *client.Client, ctx context.Context, containerId string) error {
	waiter, err := client.ContainerAttach(ctx, containerId, types.ContainerAttachOptions{
		Stderr: true,
		Stdout: true,
		Stdin:  true,
		Stream: true,
		Logs:   true,
	})

	if err != nil {
		return err
	}

	go io.Copy(os.Stdout, waiter.Reader)
	go io.Copy(os.Stderr, waiter.Reader)
	go io.Copy(waiter.Conn, os.Stdin)

	return nil
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

func RunImageWithArgs(runOptions *RunImageOptions) error {
	ctx := context.Background()

	client, err := getDefaultDockerClient()
	if err != nil {
		return err
	}

	// Pull image
	image := "638117407428.dkr.ecr.ap-south-1.amazonaws.com/cli:no-progress-bar"
	if err := PullLatestImage(image, client); err != nil {
		return err
	}

	// Generate container configurations
	containerConfig := getBaseContainerConfig(image)
	containerConfig.Cmd = runOptions.Args
	hostConfig := getContainerHostConfig(runOptions.Volumes, runOptions.Ports)

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

	// Attach input/output streams with container
	// if err := attachStdioWithContainer(client, ctx, creationResponse.ID); err != nil {
	// 	return err
	// }

	// Start container
	fmt.Printf("\n> Starting container with the latest image (ContainerId: %s)\n", creationResponse.ID)
	if err := client.ContainerStart(ctx, creationResponse.ID, types.ContainerStartOptions{}); err != nil {
		return err
	}

	// Setup interrupt fns if enabled
	if runOptions.SetupInterrupt {
		// Listen for interrupt, clear signal after execution
		// Remove container when received
		// All cleanup here: The process ends after this
		// and defer functions are not executed
		sgn := utils.RunOnCtrlC(
			func() {
				// ideally stop here, and remove later
				RemoveContainerForcefully(client, ctx, creationResponse.ID)
			},
			"Received interrupt signal",
		)
		defer utils.ClearSignals(sgn)
	}

	if runOptions.Ports.WebPortEnabled && runOptions.SpawnWebBrowser {
		quitProgressChannel := make(chan bool)
		go utils.RenderProgressSpinnerWithMessages(quitProgressChannel)
		go utils.WaitAndOpenURL(fmt.Sprintf("http://localhost:%d", runOptions.Ports.WebPortHost), quitProgressChannel, 6)
	}

	// Image output after this point
	fmt.Println("\n> Waiting for process to complete:")
	fmt.Println()

	// wait for container to stop (automatically or by interrupt)
	if err := WaitForContainer(client, ctx, creationResponse.ID); err != nil {
		return err
	}

	return nil
}
