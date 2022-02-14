package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/mitchellh/go-homedir"
)

var AppConfig *Configuration

type Configuration struct {
	HomeDirectory                 string
	DefaultWebPort                int
	ConfigurationDirectory        string
	DefaultLicensePath            string
	PrivacyResultsPathSuffix      string
	PrivacyReportsDirectorySuffix string
	PrivadoRepository             string
	Container                     *ContainerConfiguration
}

type ContainerConfiguration struct {
	ImageURL            string
	SourceCodeVolumeDir string
	LicenseVolumeDir    string
	WebPort             string
}

func init() {
	home, _ := homedir.Dir()

	imageTag := "latest"
	licenseFileName := "license.json"

	if isDev, err := strconv.ParseBool(os.Getenv("PRIVADO_DEV")); err == nil && isDev {
		imageTag = "dev"
		licenseFileName = "license-dev.json"
	}

	AppConfig = &Configuration{
		HomeDirectory:                 home,
		DefaultWebPort:                3000,
		ConfigurationDirectory:        filepath.Join(home, ".privado"),
		DefaultLicensePath:            filepath.Join(home, ".privado", licenseFileName),
		PrivacyResultsPathSuffix:      filepath.Join(".privado", "privacy.json"),
		PrivacyReportsDirectorySuffix: filepath.Join(".privado", "reports"),
		PrivadoRepository:             "https://github.com/Privado-Inc/privado",
		Container: &ContainerConfiguration{
			ImageURL:            fmt.Sprintf("public.ecr.aws/privado/cli:%s", imageTag),
			SourceCodeVolumeDir: "/app/code",
			LicenseVolumeDir:    "/tmp/license.json",
			WebPort:             "80/tcp",
		},
	}
}
