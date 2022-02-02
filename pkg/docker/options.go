package docker

type containerVolumes struct {
	licenseVolumeEnabled, sourceCodeVolumeEnabled bool
	licenseVolumeHost, sourceCodeVolumeHost       string
}

type containerPorts struct {
	webPortEnabled bool
	webPortHost    int
}

type RunImageOption func(opts *runImageHandler)

type runImageHandler struct {
	args              []string
	volumes           containerVolumes
	ports             containerPorts
	setupInterrupt    bool
	spawnWebBrowser   bool
	progressLoader    bool
	afterLoadMessages []string
	attachOutput      bool
	exitErrorMessages []string
}

func newRunImageHandler(opts []RunImageOption) runImageHandler {
	// defaults here
	rh := runImageHandler{}
	for _, opt := range opts {
		opt(&rh)
	}
	return rh
}

// Prepend option functions with "Option"

func OptionWithArgs(args []string) RunImageOption {
	return func(rh *runImageHandler) {
		rh.args = args
	}
}

func OptionWithLicenseVolume(volumeHost string) RunImageOption {
	return func(rh *runImageHandler) {
		rh.volumes.licenseVolumeEnabled = true
		rh.volumes.licenseVolumeHost = volumeHost
	}
}

func OptionWithSourceVolume(volumeHost string) RunImageOption {
	return func(rh *runImageHandler) {
		rh.volumes.sourceCodeVolumeEnabled = true
		rh.volumes.sourceCodeVolumeHost = volumeHost
	}
}

func OptionWithWebPort(portHost int) RunImageOption {
	return func(rh *runImageHandler) {
		rh.ports.webPortEnabled = true
		rh.ports.webPortHost = portHost
	}
}

func OptionWithInterrupt() RunImageOption {
	return func(rh *runImageHandler) {
		rh.setupInterrupt = true
	}
}

func OptionWithAutoSpawnBrowser() RunImageOption {
	return func(rh *runImageHandler) {
		rh.spawnWebBrowser = true
	}
}

func OptionWithProgressLoader(afterLoadMessages []string) RunImageOption {
	return func(rh *runImageHandler) {
		rh.progressLoader = true
		if len(afterLoadMessages) > 0 {
			rh.afterLoadMessages = afterLoadMessages
		}
	}
}

func OptionWithExitErrorMessages(messages []string) RunImageOption {
	return func(rh *runImageHandler) {
		rh.exitErrorMessages = messages
	}
}

func OptionWithAttachedOutput() RunImageOption {
	return func(rh *runImageHandler) {
		rh.attachOutput = true
	}
}

func OptionWithDebug(isDebug bool) RunImageOption {
	return func(rh *runImageHandler) {
		// currently only enable output in debug mode
		if isDebug {
			rh.attachOutput = true
		}
	}
}
