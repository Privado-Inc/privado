package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
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

func update(cmd *cobra.Command, args []string) {
	version(cmd, args)
	fmt.Println()

	// if Version == "dev" {
	// 	exit(
	// 		fmt.Sprint("Cannot perform an update on the dev build. Kindly use a release build or update manually\nFor more information, visit ", config.AppConfig.PrivadoRepository),
	// 		false,
	// 	)
	// }
	fmt.Println("Fetching latest release..")
	releaseInfo, err := utils.GetLatestReleaseFromGitHub(config.AppConfig.PrivadoRepositoryName)
	if releaseInfo.TagName == "" || releaseInfo.PublishedAt == "" || err != nil {
		exit("Could not fetch latest release. Some error occurred", true)
	}
	// Version = "v0.1"
	if semver.Compare(releaseInfo.TagName, Version) < 1 {
		exit(fmt.Sprint("You are already using the latest version of Privado CLI: ", Version), false)
	}

	// timeSinceRelease := int(time.Since(publishedAtTime.Local()).Hours() / 24)

	// fmt.Printf("New release found: %s (Released %d days ago)\n", releaseInfo.TagName, timeSinceRelease)
	time.Sleep(config.AppConfig.SlowdownTime)

	// create temporary directory for update assets
	temporaryDirectory, err := ioutil.TempDir("", "privado-update-")
	if err != nil {
		exit("Could not create temporary download file. Terminating..", true)
	}

	// get download url
	replacer := strings.NewReplacer(
		"${REPO_NAME}", config.AppConfig.PrivadoRepositoryName,
		"${REPO_TAG}", "latest",
		"${REPO_RELEASE_FILE}", config.AppConfig.PrivadoRepositoryReleaseFilename,
	)
	githubReleaseDownloadURL := replacer.Replace(config.ExtConfig.GitHubReleaseDownloadURL)

	// download to file
	downloadedFilePath := filepath.Join(temporaryDirectory, config.AppConfig.PrivadoRepositoryReleaseFilename)
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

	fmt.Println("githubReleaseFileURL", githubReleaseDownloadURL, temporaryDirectory)

	printVersion := Version
	if Version == "dev" {
		printVersion = "Nightly"
	}
	fmt.Printf("Privado CLI: Version %s (%s-%s) \n", printVersion, runtime.GOOS, runtime.GOARCH)
	fmt.Println("For more information, visit", config.AppConfig.PrivadoRepository)

	file, _ := os.Executable()
	fmt.Println(file)
	final, _ := filepath.EvalSymlinks(file)
	fmt.Println(final)

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
