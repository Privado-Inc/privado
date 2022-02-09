package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/Privado-Inc/privado/pkg/config"
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
	scanCmd.Flags().BoolP("overwrite", "o", false, "If specified, the warning prompt for existing scan results is disabled and any existing results are overwritten")
	scanCmd.Flags().Bool("debug", false, "Debug flag to enable image output")
	scanCmd.Flags().MarkHidden("debug")
}

func scan(cmd *cobra.Command, args []string) {
	repository := args[0]
	port, err := cmd.Flags().GetInt("port")
	debug, _ := cmd.Flags().GetBool("debug")
	overwriteResults, _ := cmd.Flags().GetBool("overwrite")
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

	// [TODO]: Check for report and prompt
	// Also add another flag to scan skip prompt
	// if overwrite flag is not specified, check for existing results
	if !overwriteResults {
		resultsPath := filepath.Join(utils.GetAbsolutePath(repository), config.AppConfig.PrivacyResultsPathSuffix)
		if exists, _ := utils.DoesFileExists(resultsPath); exists {
			fmt.Printf("> Scan report already exists (%s)\n", config.AppConfig.PrivacyResultsPathSuffix)
			fmt.Println("> If you want to view or edit existing results, run 'privado load' instead")

			fmt.Println("\n> Rescan will overwrite existing results and progress")
			confirm, _ := utils.ShowConfirmationPrompt("Continue?")
			if !confirm {
				exit("Terminating..", false)
			}
			fmt.Println()
		}
	}

	fmt.Println("> Scanning directory:", utils.GetAbsolutePath(repository))
	// run image with options
	err = docker.RunImage(
		docker.OptionWithSourceVolume(utils.GetAbsolutePath(repository)),
		docker.OptionWithLicenseVolume(utils.GetAbsolutePath(licensePath)),
		docker.OptionWithWebPort(port),
		docker.OptionWithInterrupt(),
		docker.OptionWithAutoSpawnBrowser(),
		docker.OptionWithProgressLoader(
			[]string{
				"Scanning repository..",
				"Scanning can take upto 5-10 minutes",
			},
			[]string{
				"Scanning Complete",
				fmt.Sprintf("Results and reports can be viewed on http://localhost:%d\n", port),
				"Press CTRL+C anytime to terminate the process",
				"A gentle reminder to save your changes before you exit",
			},
		),
		docker.OptionWithExitErrorMessages([]string{
			// known faults and errors
			"App: Encountered error: ",
			"Segmentation Fault",
			"Some error occurred while processing analyzer results!",
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
	defineScanFlags(scanCmd)
	rootCmd.AddCommand(scanCmd)
}
