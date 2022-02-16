package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/Privado-Inc/privado/pkg/config"
	"github.com/Privado-Inc/privado/pkg/docker"
	"github.com/Privado-Inc/privado/pkg/utils"
	"github.com/spf13/cobra"
)

var loadCmd = &cobra.Command{
	Use:   "load <repository>",
	Short: "Load a scanned codebase to continue generating compliance reports",
	Long:  "Load a scanned codebase or repository and continue generating compliance reports. It skips privacy scan and loads the results present in the target repository (.privado directory)",
	Args:  cobra.ExactArgs(1),
	Run:   load,
}

func defineLoadFlags(cmd *cobra.Command) {
	loadCmd.Flags().IntP("port", "p", 3000, "The port to be used to render HTML results")
	loadCmd.Flags().Bool("debug", false, "Debug flag to enable image output")
	loadCmd.Flags().MarkHidden("debug")
}

func load(cmd *cobra.Command, args []string) {
	repository := args[0]
	port, err := cmd.Flags().GetInt("port")
	debug, _ := cmd.Flags().GetBool("debug")
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

	fmt.Println("> Loading results from directory:", utils.GetAbsolutePath(repository))
	resultsPath := filepath.Join(utils.GetAbsolutePath(repository), config.AppConfig.PrivacyResultsPathSuffix)
	if exists, _ := utils.DoesFileExists(resultsPath); !exists {
		exit("Cannot find scan results for the specified directory.\nRun 'privado scan <dir>' instead.\n\nRun 'privado help' for more information.", true)
	}

	// run image with options
	err = docker.RunImage(
		docker.OptionWithArgs([]string{"--no-scan"}),
		docker.OptionWithSourceVolume(utils.GetAbsolutePath(repository)),
		docker.OptionWithLicenseVolume(utils.GetAbsolutePath(licensePath)),
		docker.OptionWithWebPort(port),
		docker.OptionWithInterrupt(),
		docker.OptionWithAutoSpawnBrowser(),
		docker.OptionWithProgressLoader(
			[]string{"Loading results.."},
			[]string{
				fmt.Sprintf("Results and reports can be viewed on http://localhost:%d\n", port),
				"Press CTRL+C anytime to terminate the process",
				"A gentle reminder to save your changes before you exit",
			},
		),
		docker.OptionWithExitErrorMessages([]string{
			// known faults and errors
			"App: Encountered error: ",
			"Segmentation Fault",
			"exception: ",
			"privado: error: invalid choice",
			"requests.exceptions.ConnectionError: HTTPSConnectionPool(",
			"Failed to execute script 'app' due to unhandled exception!",
		}),
		docker.OptionWithDebug(debug),
	)
	if err != nil {
		exit(fmt.Sprintf("Received error: %s", err), true)
	}
}

func init() {
	defineLoadFlags(loadCmd)
	rootCmd.AddCommand(loadCmd)
}
