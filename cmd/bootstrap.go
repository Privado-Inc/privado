package cmd

import (
	"fmt"
	"os"

	"github.com/Privado-Inc/privado/pkg/config"
	"github.com/Privado-Inc/privado/pkg/utils"
	"github.com/spf13/cobra"
)

var bootstrapCmd = &cobra.Command{
	Use:   "bootstrap <license.json>",
	Short: "Authenticates Privado using the requested license and generates required configurations",
	Long:  "Authenticates Privado using the requested license and generates required configurations. To request a license, run `privado auth <email>`.",
	Args:  cobra.ExactArgs(1),
	Run:   bootstrap,
}

func defineBootstrapFlags(cmd *cobra.Command) {
	bootstrapCmd.Flags().Bool("overwrite", false, "Overwrites the existing license file (if any)")
}

func bootstrap(cmd *cobra.Command, args []string) {
	fmt.Println("> Bootstrapping Privado CLI")

	pathToLicenseFile := args[0]
	overwriteFlag, err := cmd.Flags().GetBool("overwrite")
	if err != nil {
		exit(fmt.Sprint("Cannot parse flag --overwrite", err), true)
	}

	// verify license file format
	if err := utils.VerifyLicenseJSON(pathToLicenseFile); err != nil {
		exit(fmt.Sprintf("Could not verify license file: %s", err), true)
	}

	pathToLicenseFile = utils.GetAbsolutePath(pathToLicenseFile)

	// soft-error in case of existing license
	if fileExists, _ := utils.DoesFileExists(config.AppConfig.DefaultLicensePath); fileExists {
		if !overwriteFlag {
			exit("License already exists. Use '--overwrite' flag to replace the existing license", true)
		} else {
			fmt.Println("> '--overwrite' flag found, replacing existing license")
		}
	}

	// check if homedir was found
	if config.AppConfig.HomeDirectory == "" {
		exit("Could not detect a home directory. Unable to bootstrap", true)
	}

	// create configuration directory
	if err := os.MkdirAll(config.AppConfig.ConfigurationDirectory, os.ModePerm); err != nil {
		exit(fmt.Sprintf("Could not create configuration directory: %s", err), true)
	}

	// copy license to target location
	targetLicensePath := config.AppConfig.DefaultLicensePath
	if err := utils.CopyFile(pathToLicenseFile, targetLicensePath); err != nil {
		exit(fmt.Sprintf("Could not copy license file: %s", err), true)
	}

	// set appropriate license permissions
	if err := os.Chmod(targetLicensePath, 0600); err != nil {
		exit(fmt.Sprintf("Could not set appropriate permissions for the license file: %s", err), true)
	}

	exit("> Privado CLI bootstrapped successfully!", false)
}

func init() {
	defineBootstrapFlags(bootstrapCmd)
	rootCmd.AddCommand(bootstrapCmd)
}
