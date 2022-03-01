package cmd

import (
	"fmt"
	"runtime"
	"time"

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

	// Additional info for exclusively this cmd (so 'version' can be called just to print version)
	if cmd.Name() == "version" {
		hasUpdate, updateMessage, err := checkForUpdate()
		if err == nil && hasUpdate {
			fmt.Println()
			fmt.Println(updateMessage)
			time.Sleep(config.AppConfig.SlowdownTime)
			fmt.Println("To use the latest version of Privado CLI, run `privado update`")
			fmt.Println()
		}

		time.Sleep(config.AppConfig.SlowdownTime)
		fmt.Println("For more information, visit", config.AppConfig.PrivadoRepository)
	}
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
