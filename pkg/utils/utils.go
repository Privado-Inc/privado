package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"runtime"
	"syscall"
)

type LicenseFileType struct {
	PrivateKey string `json:"pri_key"`
	UserHash   string `json:"user_hash"`
}

func GetAbsolutePath(relativePath string) string {

	fullPath, err := filepath.Abs(relativePath)
	if err != nil {
		panic(err)
	}
	return fullPath
}

func VerifyLicenseJSON(pathToFile string) error {
	// open file
	file, err := os.Open(pathToFile)
	if err != nil {
		return err
	}
	defer file.Close()

	// read file
	dataBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	// parse file
	var license LicenseFileType
	if err := json.Unmarshal(dataBytes, &license); err != nil {
		return err
	}

	// verify license info
	if license.PrivateKey == "" || license.UserHash == "" {
		return fmt.Errorf("malformed file data")
	}

	return nil
}

func OpenURLInBrowser(url string) {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("xdg-open", url)
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	case "darwin":
		cmd = exec.Command("open", url)
	default:
		cmd = nil
	}

	if cmd != nil {
		if err := cmd.Start(); err == nil {
			return
		}
	}

	// in case we cannot automatically open due to
	// unknown OS or an error, print
	fmt.Println("\n> Unable to open browser")
	fmt.Println("> Open the following URL to view results:", url)
}

func RunOnCtrlC(cleanupFn func()) chan os.Signal {
	notifySignal := make(chan os.Signal)
	signal.Notify(notifySignal, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-notifySignal
		fmt.Println("\n> Received interrupt signal")
		cleanupFn()
		os.Exit(0)
	}()

	return notifySignal
}

func ClearSignals(sgn chan os.Signal) {
	signal.Stop(sgn)
}
