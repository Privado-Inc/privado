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
	Short: "Scan a codebase or repository to identify privacy issues and generate compliance reports",
	Args:  cobra.ExactArgs(1),
	Run:   scan,
}

func defineScanFlags(cmd *cobra.Command) {
	scanCmd.Flags().IntP("port", "p", 3000, "The port to be used to render HTML results.")
}

func scan(cmd *cobra.Command, args []string) {
	repository := args[0]
	port, err := cmd.Flags().GetInt("port")
	if err != nil {
		exit(fmt.Sprint("Cannot parse flag --port", err), true)
	}

	home, err := homedir.Dir()
	if err != nil {
		exit(fmt.Sprint("Cannot determine license directory (~/.privado)", err), true)
	}
	licensePath := path.Join(home, ".privado", "license.json")

	err = docker.RunImageWithArgs(
		docker.OptionWithSourceVolume(utils.GetAbsolutePath(repository)),
		docker.OptionWithLicenseVolume(utils.GetAbsolutePath(licensePath)),
		docker.OptionWithWebPort(port),
		docker.OptionWithInterrupt(),
		docker.OptionWithAutoSpawnBrowser(),
		docker.OptionWithProgressLoader(),
	)
	if err != nil {
		exit(fmt.Sprintf("Received error: %s", err), true)
	}
}

func init() {
	defineScanFlags(scanCmd)
	rootCmd.AddCommand(scanCmd)
}
