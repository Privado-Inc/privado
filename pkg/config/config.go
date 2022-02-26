package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"time"

	"github.com/mitchellh/go-homedir"
)

var AppConfig *Configuration
var ExtConfig *ExternalConfiguration

type Configuration struct {
	HomeDirectory                    string
	DefaultWebPort                   int
	ConfigurationDirectory           string
	DefaultLicensePath               string
	PrivacyResultsPathSuffix         string
	PrivacyReportsDirectorySuffix    string
	PrivadoRepository                string
	PrivadoRepositoryName            string
	PrivadoRepositoryReleaseFilename string
	SlowdownTime                     time.Duration
	Container                        *ContainerConfiguration
}

type ContainerConfiguration struct {
	ImageURL            string
	SourceCodeVolumeDir string
	LicenseVolumeDir    string
	WebPort             string
}

type ExternalConfiguration struct {
	GitHubAPIHost            string
	GitHubReleasesEndpoint   string
	GitHubReleaseDownloadURL string
}

// init function for AppConfig
func init() {
	home, _ := homedir.Dir()

	imageTag := "latest"
	licenseFileName := "license.json"

	if isDev, err := strconv.ParseBool(os.Getenv("PRIVADO_DEV")); err == nil && isDev {
		imageTag = os.Getenv("PRIVADO_TAG")
		if imageTag == "" {
			imageTag = "dev"
		}
		licenseFileName = "license-dev.json"
	}

	AppConfig = &Configuration{
		HomeDirectory:                    home,
		DefaultWebPort:                   3000,
		ConfigurationDirectory:           filepath.Join(home, ".privado"),
		DefaultLicensePath:               filepath.Join(home, ".privado", licenseFileName),
		PrivacyResultsPathSuffix:         filepath.Join(".privado", "privacy.json"),
		PrivacyReportsDirectorySuffix:    filepath.Join(".privado", "reports"),
		PrivadoRepository:                "https://github.com/Privado-Inc/privado",
		PrivadoRepositoryName:            "Privado-Inc/privado",
		PrivadoRepositoryReleaseFilename: fmt.Sprintf("privado-%s-%s.tar.gz", runtime.GOOS, runtime.GOARCH),
		SlowdownTime:                     600 * time.Millisecond,
		Container: &ContainerConfiguration{
			ImageURL:            fmt.Sprintf("public.ecr.aws/privado/cli:%s", imageTag),
			SourceCodeVolumeDir: "/app/code",
			LicenseVolumeDir:    "/tmp/license.json",
			WebPort:             "80/tcp",
		},
	}
}

// init function for ExtConfig
func init() {
	ExtConfig = &ExternalConfiguration{
		GitHubAPIHost:            "https://api.github.com",
		GitHubReleasesEndpoint:   "/repos/${REPO_NAME}/releases/latest",
		GitHubReleaseDownloadURL: "https://github.com/${REPO_NAME}/releases/download/${REPO_TAG}/${REPO_RELEASE_FILE}",
	}
}
