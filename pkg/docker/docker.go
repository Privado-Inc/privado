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
	Args           []string
	Volumes        *ContainerVolumes
	Ports          *ContainerPorts
	SetupInterrupt bool
}

func getDefaultDockerClient() (*client.Client, error) {
	client, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}

	return client, nil
}

func getBaseContainerConfig() *container.Config {
	config := &container.Config{
		Image:        "638117407428.dkr.ecr.ap-south-1.amazonaws.com/cli:latest",
		AttachStdin:  true,
		AttachStdout: true,
		AttachStderr: true,
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

// 1. pull image
// 2. create container
// 3. run container
// handle o/p i/p

func PullLatestImage(image string, client *client.Client) (err error) {
	if client == nil {
		client, err = getDefaultDockerClient()
		if err != nil {
			return err
		}
	}

	auth := types.AuthConfig{
		Username:      "AWS",
		Password:      "eyJwYXlsb2FkIjoiS2dJOTFVTWxUcXlWdE9qTEhLenlBMkJBSnRKRCtoV0pRQ2hZOC8rMys3OFl3TTZ0K1hjMFhZOWdBenBwV0tXcUt6MVdOMytnQ0dCSnlDUUlpc3FPcGVaVE8xbFlVMWltTFhWY1M2dis5VG9SMUNvclhRR2VBVTJaZTUxSjl3M0grUE1FUjAvMDZSczlrczhVVzdFRXhUZTg0UmFjK3BSS1UzMGFNbndoM3ljUThsaWE2TzdKdFZ4eDZYR2U0TE5kMm12VzhkWGFpQktBTjJSYyswV1ljSE8yWWtEclVEZHBpcXRzMTFzVXRkQ3VkVmtRKzB0LzZ5ZVB6b1NwWEVZWkZHREcrY1JQV1B1RHgyeHY4b0I3aVltZm9ucjRiWm9ZbXcreHdNTTlDQWEzK3U1TEE0OXZ5b21FTkt4UHJaTTJkWE5zdWd2TTNmbDVsMUFNTUZ6M014dW5XWnlhUkJPeGN2S1YxM0lMM2ZwMEU1Mm9PRTg2aVh3UTdha2Jsd0RtRWZMTlZxcWNWUmFNREM2ZE5mUGRhc1dKOEgwa2REVzJ1U1ZzbFBlWnhXbVZjRUdneWZzYlYvTTNZMTZoTjY0RWpYd1ZENHFWUE1qREx2bHhiWGcxK3Q4ODAvbmNjQU9TbjJLZktacXZEbm1zUEVyN2lKUVBPUGFNNUVrcUtSWlhIYTZWK1FBdmFyL3NvSVROL2RRY0hhQldDdkpyWjhjWFh4ekxjZFFGWHZRNXg0MHhQak53RlJpRzNhMGZ5T216TzBKWWtQcU1IWnduN0dPSWNkWVlCV25nekR4RkNycVUzM2NhV2tvRlJSaXhNUnkxU1I1MU1LTjBXY3ZNcFFqUU0wWFJ5SFQzVEFVaVpNYlY4VEoxSjE4WDE3Yk1WVWpPeVhPTzhlL3laT3hYYklVc0pXUWZLbXdKa0dwOHI1dnI2M0N6SitwV3hPaFZsdU9PTGZUYnYwbzNINnc1ZmtQcGV5ZE4xNm5RWXd1QTNJZzdnbjdZdGgzZC9UYWtzTmlhN2tieG9oaEd6VW55d2hLNTkvRTl2YjVrS0lFci9IYzJyb0VMSllpYmo5U1N3eWQyTW1WNDlOaFBjWFVLdnE0NWVKeFZzb2IvSUR5ekdsYWEwNEZmcVFISUF6WlBZMERWb213OFFOMXBpSXZtVVhKSUF0N2FrditZWTV3NnlVamxhV2lndVRIcjl4VHpmbnpaTFZ1SWpYYThFNkRVSEFoZnRndktBaStCUFJ2Y1BGZUIxbHBUQTMvVDJTV01wQWNWdFF6K09UQ3hnVVpraWwzSnJDQ256QmJYUjZjUDd5QndJZ1JMeS9iOFcvUi9nUjhYL2NVZHFRS0VOY1laRnJ4c2gyR1R3YlloQXFPODY0YVFWYjU3V0Y2LzJ1Z0ZJN3VvdjJRcCtNcz0iLCJkYXRha2V5IjoiQVFJQkFIaUhXYVlUblJVV0NibnorN0x2TUcrQVB2VEh6SGxCVVE5RnFFbVYyNkJkd3dIRXBtQnRFY3BrcExaaFZQSGFtTmx1QUFBQWZqQjhCZ2txaGtpRzl3MEJCd2FnYnpCdEFnRUFNR2dHQ1NxR1NJYjNEUUVIQVRBZUJnbGdoa2dCWlFNRUFTNHdFUVFNaWVnK2VCOXY5YXJqM0hrS0FnRVFnRHRoa1BJNmdkaW5JOGRaL0I4MzJyUzE4Vy9JUkEwSVc5OS9ZMjRWWXh4TTZqU0ZJbGpFenQrZ2lDZFpneEhXZ2lRaGdiMFBWb0NadVRGWHFBPT0iLCJ2ZXJzaW9uIjoiMiIsInR5cGUiOiJEQVRBX0tFWSIsImV4cGlyYXRpb24iOjE2NDMzMDc2Mjl9",
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

	defer reader.Close()
	io.Copy(os.Stdout, reader)

	return nil
}

func attachStdioWithContainer(client *client.Client, ctx context.Context, containerId string) error {
	waiter, err := client.ContainerAttach(ctx, containerId, types.ContainerAttachOptions{
		Stderr: true,
		Stdout: true,
		Stdin:  true,
		Stream: true,
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
	if err := PullLatestImage("638117407428.dkr.ecr.ap-south-1.amazonaws.com/cli:latest", client); err != nil {
		return err
	}

	// Generate container configurations
	containerConfig := getBaseContainerConfig()
	containerConfig.Cmd = runOptions.Args
	hostConfig := getContainerHostConfig(runOptions.Volumes, runOptions.Ports)

	// Create container
	creationResponse, err := client.ContainerCreate(ctx, containerConfig, hostConfig, nil, nil, "")
	if err != nil {
		return err
	}

	// Attach input/output streams with container
	if err := attachStdioWithContainer(client, ctx, creationResponse.ID); err != nil {
		return err
	}

	// Start container
	fmt.Println("\n> Starting container with the latest image", creationResponse.ID)
	if err := client.ContainerStart(ctx, creationResponse.ID, types.ContainerStartOptions{}); err != nil {
		return err
	}

	// Image output after this point
	fmt.Println("\n> Waiting for process to complete:")
	fmt.Println()

	// Setup interrupt if enabled
	if runOptions.SetupInterrupt {
		// Listen for interrupt, clear signal after execution
		// Stop container when received
		sgn := utils.RunOnCtrlC(
			func() {
				StopContainer(client, ctx, creationResponse.ID)
			},
		)
		defer utils.ClearSignals(sgn)
	}

	// wait for container to stop (automatically or by interrupt)
	if err := WaitForContainer(client, ctx, creationResponse.ID); err != nil {
		return err
	}

	// remove the container (in all cases)
	if err := RemoveContainerForcefully(client, ctx, creationResponse.ID); err != nil {
		return err
	}

	return nil
}
