package cmd

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"time"

	"github.com/Privado-Inc/privado/pkg/config"
	"github.com/Privado-Inc/privado/pkg/utils"
	"github.com/spf13/cobra"
	"golang.org/x/mod/semver"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Check for latest release and update to the latest version Privado CLI",
	Long:  "Check for latest release and update to the latest version Privado CLI",
	Args:  cobra.ExactArgs(0),
	Run:   update,
}

func checkForUpdate() (hasUpdate bool, updateMessage string, err error) {
	if Version == "dev" {
		return false, "", nil
	}

	// get release info
	releaseInfo, err := utils.GetLatestReleaseFromGitHub(config.AppConfig.PrivadoRepositoryName)
	if releaseInfo.TagName == "" || releaseInfo.PublishedAt == "" || err != nil {
		return false, "", err
	}

	// compare release, -1, 0, 1
	if semver.Compare(releaseInfo.TagName, Version) > 0 {
		hasUpdate = true
		// Get new release information (with time elapsed if possible)
		daysSinceRelease, err := utils.GetDaysSinceRFC3339String(releaseInfo.PublishedAt)
		if err != nil {
			updateMessage = fmt.Sprintf("New release found: %s\n", releaseInfo.TagName)
		} else {
			daySinceString := ""
			switch {
			case (daysSinceRelease < 1):
				daySinceString = "Released today"
			case daysSinceRelease == 1:
				daySinceString = "Released yesterday"
			default:
				daySinceString = fmt.Sprintf("Released %d days ago", daysSinceRelease)
			}
			updateMessage = fmt.Sprintf("New release found: %s (%s)", releaseInfo.TagName, daySinceString)
		}
	}

	return hasUpdate, updateMessage, nil
}

func update(cmd *cobra.Command, args []string) {
	version(cmd, args)
	fmt.Println()
	time.Sleep(config.AppConfig.SlowdownTime)
	// if Version == "dev" {
	// 	exit(
	// 		fmt.Sprint("Cannot perform an update on the dev build. Kindly use a release build or update manually\nFor more information, visit ", config.AppConfig.PrivadoRepository),
	// 		false,
	// 	)
	// }

	// get path to current executable
	currentExecPath, err := utils.GetPathToCurrentBinary()
	if err != nil {
		exit(fmt.Sprint("Could not evaluate path to current binary. Auto update fail\nFor more information, visit", config.AppConfig.PrivadoRepository), true)
	}
	// check for appropriate permissions
	// hasPerm := utils.HasPermissionToFile("/usr/local/bin/syringe")
	// fmt.Println("Permission to write", hasPerm)
	// exit("", false)

	Version = "v0.1.1"
	// check for release info
	fmt.Println("Fetching latest release..")
	hasUpdate, updateMessage, err := checkForUpdate()
	if err != nil {
		exit("Could not fetch latest release. Some error occurred", true)
	}
	if !hasUpdate {
		exit(fmt.Sprint("You are already using the latest version of Privado CLI: ", Version), false)
	}
	fmt.Println(updateMessage)
	time.Sleep(config.AppConfig.SlowdownTime)

	// get download url
	replacer := strings.NewReplacer(
		"${REPO_NAME}", config.AppConfig.PrivadoRepositoryName,
		"${REPO_TAG}", "latest",
		"${REPO_RELEASE_FILE}", config.AppConfig.PrivadoRepositoryReleaseFilename,
	)
	githubReleaseDownloadURL := replacer.Replace(config.ExtConfig.GitHubReleaseDownloadURL)

	// create temporary directory for update assets
	temporaryDirectory, err := ioutil.TempDir("", "privado-update-")
	if err != nil {
		exit("Could not create temporary download file. Terminating..", true)
	}
	downloadedFilePath := filepath.Join(temporaryDirectory, config.AppConfig.PrivadoRepositoryReleaseFilename)

	// download to file in temp directory
	err = utils.DownloadToFile(githubReleaseDownloadURL, downloadedFilePath)
	if err != nil {
		exit(fmt.Sprint("Could not download release asset: ", githubReleaseDownloadURL), true)
	}
	fmt.Println("Downloaded release asset:", githubReleaseDownloadURL)
	time.Sleep(config.AppConfig.SlowdownTime)

	// extract .tar.gz
	fmt.Println()
	fmt.Println("Extracting release asset..")
	err = utils.ExtractTarGzFile(downloadedFilePath, temporaryDirectory)
	if err != nil {
		exit(fmt.Sprint("Could not extract release asset: ", githubReleaseDownloadURL, err), true)
	}

	fmt.Println("Extracted release asset:", temporaryDirectory)
	time.Sleep(config.AppConfig.SlowdownTime)
	fmt.Println()

	// Replace existing binary (in current execution) by the updated binary
	fmt.Println("Installing latest release..")
	time.Sleep(config.AppConfig.SlowdownTime)
	err = utils.SafeMoveFile(filepath.Join(temporaryDirectory, "privado"), currentExecPath, true)
	if err != nil {
		exit(fmt.Sprint("Could not update existing installation: ", err), true)
	}

	// woof! all done.
	time.Sleep(config.AppConfig.SlowdownTime)
	fmt.Println()
	fmt.Println("Installed latest release!")
	exit("To validate installation, run `privado version`", false)
}

func init() {
	rootCmd.AddCommand(updateCmd)
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
