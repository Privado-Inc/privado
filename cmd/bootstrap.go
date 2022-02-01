package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Privado-Inc/privado/pkg/utils"
	"github.com/mitchellh/go-homedir"
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

	if fileExists, _ := utils.DoesFileExists(pathToLicenseFile); fileExists {
		if !overwriteFlag {
			exit("License already exists. Use '--overwrite' flag to replace the existing license", true)
		}
	}

	// create configuration directory
	home, err := homedir.Dir()
	if err != nil {
		exit(fmt.Sprintf("Could not detect a home directory. Unable to bootstrap: %s", err), true)
	}
	configurationDirectory := filepath.Join(home, ".privado")
	if err := os.MkdirAll(configurationDirectory, os.ModePerm); err != nil {
		exit(fmt.Sprintf("Could not create configuration directory: %s", err), true)
	}

	targetLicensePath := filepath.Join(home, ".privado", "license.json")
	if err := utils.CopyFile(pathToLicenseFile, targetLicensePath); err != nil {
		exit(fmt.Sprintf("Could not copy license file: %s", err), true)
	}

	if err := os.Chmod(targetLicensePath, 0600); err != nil {
		exit(fmt.Sprintf("Could not set appropriate permissions for the license file: %s", err), true)
	}

	exit("> Privado CLI bootstrapped successfully!", false)
}

func init() {
	defineBootstrapFlags(bootstrapCmd)
	rootCmd.AddCommand(bootstrapCmd)
}
