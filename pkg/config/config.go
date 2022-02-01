package config

import (
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

var AppConfig *Configuration

type Configuration struct {
	HomeDirectory          string
	DefaultWebPort         int
	ConfigurationDirectory string
	DefaultLicensePath     string
	PrivadoRepository      string
	Container              *ContainerConfiguration
}

type ContainerConfiguration struct {
	ImageURL            string
	SourceCodeVolumeDir string
	LicenseVolumeDir    string
	WebPort             string
}

func init() {
	home, _ := homedir.Dir()

	AppConfig = &Configuration{
		HomeDirectory:          home,
		DefaultWebPort:         3000,
		ConfigurationDirectory: filepath.Join(home, ".privado"),
		DefaultLicensePath:     filepath.Join(home, ".privado", "license.json"),
		PrivadoRepository:      "https://github.com/Privado-Inc/privado",
		Container: &ContainerConfiguration{
			ImageURL:            "638117407428.dkr.ecr.ap-south-1.amazonaws.com/cli:no-progress-bar",
			SourceCodeVolumeDir: "/app/code",
			LicenseVolumeDir:    "/tmp/license.json",
			WebPort:             "80/tcp",
		},
	}
}
