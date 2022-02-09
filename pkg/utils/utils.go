package utils

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
	"time"

	"github.com/schollz/progressbar/v3"
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

func CopyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}

func DoesFileExists(name string) (bool, error) {
	_, err := os.Stat(name)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
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

// Ignores all error and waits for URL to be responsive
// by sending HEAD request every intervalSeconds
func WaitForResponsiveURL(url string, intervalSeconds int) {
	if intervalSeconds == 0 {
		intervalSeconds = 5
	}
	ticker := time.NewTicker(time.Duration(intervalSeconds) * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		res, err := http.Head(url)
		if err == nil && res.StatusCode == 200 {
			return
		}
	}
}

func WaitAndOpenURL(url string, sgn chan bool, interval int) {
	WaitForResponsiveURL(url, interval)
	sgn <- true
	OpenURLInBrowser(url)
}

func RunOnCtrlC(cleanupFn func()) chan os.Signal {
	notifySignal := make(chan os.Signal)
	signal.Notify(notifySignal, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-notifySignal
		cleanupFn()
		os.Exit(0)
	}()

	return notifySignal
}

func ClearSignals(sgn chan os.Signal) {
	signal.Stop(sgn)
}

func RenderProgressSpinnerWithMessages(complete, quit chan bool, loadMessages, afterLoadMessages []string) {
	if len(loadMessages) == 0 {
		// default message
		loadMessages = []string{"Loading.."}
	}

	messageIndex := 0
	messageRotationTicker := time.NewTicker(20 * time.Second)

	bar := progressbar.NewOptions(-1,
		progressbar.OptionFullWidth(),
		// Good spinners: 0, 31, 51, 52, 54
		progressbar.OptionSpinnerType(52),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionSetDescription(loadMessages[messageIndex]),
		progressbar.OptionShowCount(),
	)

	for {
		seconds := bar.State().SecondsSince
		select {
		case <-quit:
			fmt.Println()
			bar.Close()
		case <-complete:
			bar.Close()
			fmt.Println()
			fmt.Println("> Complete")
			fmt.Println("> Total Time taken:", seconds, "seconds")

			if len(afterLoadMessages) > 0 {
				for _, message := range afterLoadMessages {
					fmt.Print("\n> ", message)
				}
			}
			fmt.Println()
			return
		case <-messageRotationTicker.C:
			messageIndex++
			bar.Describe(loadMessages[messageIndex%len(loadMessages)])
		default:
			bar.Set(int(seconds))
			time.Sleep(150 * time.Millisecond)
		}
	}
}

func ShowConfirmationPrompt(msg string) (bool, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s (y/N): ", msg)
	ans, err := reader.ReadString('\n')
	if err != nil {
		return false, err
	}
	ans = strings.TrimSpace(ans)
	ans = strings.ToLower(ans)

	if ans == "y" || ans == "yes" || ans == "1" {
		return true, nil
	}
	return false, nil
}
