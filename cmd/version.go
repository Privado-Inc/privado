package cmd

import (
	"fmt"
	"runtime"

	"github.com/Privado-Inc/privado/pkg/config"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the current version of Privado CLI",
	Long:  "Print the current version of Privado CLI",
	Args:  cobra.ExactArgs(0),
	Run:   version,
}

func version(cmd *cobra.Command, args []string) {
	printVersion := Version
	if Version == "dev" {
		printVersion = "Nightly"
	}
	fmt.Printf("Privado CLI: Version %s (%s-%s) \n", printVersion, runtime.GOOS, runtime.GOARCH)
	fmt.Println("For more information, visit", config.AppConfig.PrivadoRepository)
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

// update logic
// Check for latest version tag
// Get the latest tag
// Compare if new available
// Prompt user
// if yes: download binary for respective os & arch
// Get location of current binary
// for UNIX:
// Check if permissions available to replace the binary
// If yes, replace
// If not, prompt? or what? sudo ?? > Prompt with command
// for non-UNIX i.e. windows:
// not possible to replace files during runtime
// prompt command
