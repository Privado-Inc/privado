package cmd

import (
	"fmt"

	"github.com/Privado-Inc/privado/pkg/docker"
	"github.com/Privado-Inc/privado/pkg/utils"
	"github.com/spf13/cobra"
)

var scanCmd = &cobra.Command{
	Use:   "scan <repository>",
	Short: "Scan a codebase or repository to identify privacy issues and generate compliance reports",
	Args:  cobra.ExactArgs(1),
	Run:   scan,
}

func defineScanFlags(cmd *cobra.Command) {
	scanCmd.Flags().IntP("port", "p", 3000, "The port to be used to render HTML results")
	scanCmd.Flags().Bool("debug", false, "Debug flag to enable image output")
	scanCmd.Flags().MarkHidden("debug")
}

func scan(cmd *cobra.Command, args []string) {
	repository := args[0]
	port, err := cmd.Flags().GetInt("port")
	if err != nil {
		exit(fmt.Sprint("Cannot parse flag --port", err), true)
	}

	// verify license
	if err := utils.VerifyLicenseJSON(licensePath); err != nil {
		exit(fmt.Sprint(
			fmt.Sprintf("Could not verify license: %s\n", err),
			"To request a license, run: 'privado auth <email>'\n",
			"To bootstrap the app, run: 'privado bootstrap <license>'\n\n",
			"For more info, run: 'privado help'\n",
		), true)
	}

	fmt.Println("> Scanning directory:", utils.GetAbsolutePath(repository))

	// run image with options
	err = docker.RunImageWithArgs(
		docker.OptionWithSourceVolume(utils.GetAbsolutePath(repository)),
		docker.OptionWithLicenseVolume(utils.GetAbsolutePath(licensePath)),
		docker.OptionWithWebPort(port),
		docker.OptionWithInterrupt(),
		docker.OptionWithAutoSpawnBrowser(),
		docker.OptionWithProgressLoader([]string{
			fmt.Sprintf("Results and reports can be viewed on http://localhost:%d\n", port),
			"Press CTRL+C anytime to terminate the process",
			"A gentle reminder to save your changes before you exit",
		}),
		docker.OptionWithExitErrorMessages([]string{
			// known faults and errors
			"Segmentation Fault",
			"Some error occurred while processing analyzer results!",
			"exception: ",
			"privado: error: invalid choice",
			"requests.exceptions.ConnectionError: HTTPSConnectionPool(",
			"Failed to execute script 'app' due to unhandled exception!",
		}),
	)
	if err != nil {
		exit(fmt.Sprintf("Received error: %s", err), true)
	}
}

func init() {
	defineScanFlags(scanCmd)
	rootCmd.AddCommand(scanCmd)
}
