package cmd

import (
	"fmt"
	"path"

	"github.com/Privado-Inc/privado/pkg/docker"
	"github.com/Privado-Inc/privado/pkg/utils"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

var scanCmd = &cobra.Command{
	Use:   "scan <repository>",
	Short: "Request license for Privado CLI",
	Args:  cobra.ExactArgs(1),
	Run:   scan,
}

func defineScanFlags(cmd *cobra.Command) {
	scanCmd.Flags().IntP("port", "p", 3000, "The port to be used to render HTML results.")
}

func scan(cmd *cobra.Command, args []string) {
	repository := args[0]

	home, err := homedir.Dir()
	if err != nil {
		exit("Cannot determine license directory (~/.privado)", true)
	}
	licensePath := path.Join(home, ".privado", "license.json")

	runImageOptions := &docker.RunImageOptions{
		SetupInterrupt: true, // Setup stop on Ctrl+C
		Volumes: &docker.ContainerVolumes{
			LicenseVolumeEnabled:    true,
			LicenseVolumeHost:       utils.GetAbsolutePath(licensePath),
			SourceCodeVolumeEnabled: true,
			SourceCodeVolumeHost:    utils.GetAbsolutePath(repository),
		},
		Ports: &docker.ContainerPorts{
			WebPortEnabled: true,
			WebPortHost:    3000,
		},
	}

	err = docker.RunImageWithArgs(runImageOptions)
	if err != nil {
		exit(fmt.Sprintf("Received error: %s", err), true)
	}

}

func init() {
	defineScanFlags(scanCmd)
	rootCmd.AddCommand(scanCmd)
}
