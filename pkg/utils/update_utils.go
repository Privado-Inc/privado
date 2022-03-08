package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Privado-Inc/privado/pkg/config"
	extract "github.com/codeclysm/extract/v3"
	"github.com/schollz/progressbar/v3"
)

type gitHubReleaseType struct {
	TagName     string `json:"tag_name"`
	PublishedAt string `json:"published_at"`
}

func GetLatestReleaseFromGitHub(repoName string) (*gitHubReleaseType, error) {
	endpoint := strings.Replace(config.ExtConfig.GitHubReleasesEndpoint, "${REPO_NAME}", config.AppConfig.PrivadoRepositoryName, 1)
	url := fmt.Sprintf("%s%s", config.ExtConfig.GitHubAPIHost, endpoint)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	defaultClient := &http.Client{}
	response, err := defaultClient.Do(req)

	if err != nil || response.StatusCode != 200 {
		return nil, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	releaseResponse := gitHubReleaseType{}
	err = json.Unmarshal(responseData, &releaseResponse)
	if err != nil {
		return nil, err
	}

	return &releaseResponse, nil
}

func DownloadToFile(downloadURL, filePath string) error {
	req, err := http.NewRequest("GET", downloadURL, nil)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	bar := progressbar.DefaultBytes(
		resp.ContentLength,
		"Downloading..",
	)

	_, err = io.Copy(io.MultiWriter(file, bar), resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func ExtractTarGzFile(sourceFile, target string) error {
	ctx := context.Background()
	data, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		return err
	}
	buffer := bytes.NewBuffer(data)
	err = extract.Gz(ctx, buffer, target, func(s string) string { return s })
	if err != nil {
		return err
	}

	return nil
}

func GetDaysSinceRFC3339String(date string) (int, error) {
	parsedDate, err := time.Parse(time.RFC3339, date)
	if err != nil {
		return 0, err
	}
	elapsedTime := time.Since(parsedDate)
	elapsedDays := int(math.Round(elapsedTime.Hours() / 24))
	return elapsedDays, nil
}

func GetPathToCurrentBinary() (string, error) {
	currentFilePath, err := os.Executable()
	if err != nil {
		return "", err
	}
	// resolve any symlinks
	resolvedFilePath, err := filepath.EvalSymlinks(currentFilePath)
	if err != nil {
		return "", err
	}

	return resolvedFilePath, nil
}

func HasWritePermissionToFile(filePath string) (bool, error) {
	file, err := os.OpenFile(filePath, os.O_RDWR, 0744)
	if err != nil {
		if errors.Is(err, fs.ErrPermission) {
			return false, nil
		}
		return false, err
	}
	defer file.Close()

	return true, nil
}

func SafeMoveFile(source, target string, showLogs bool) (err error) {
	// Resolve symlinks and use actual path for critical operation
	if source, err = filepath.EvalSymlinks(source); err != nil {
		return err
	}
	if target, err = filepath.EvalSymlinks(target); err != nil {
		return err
	}

	fileExists, err := DoesFileExists(target)
	if err != nil {
		return err
	}

	// if file exists, make a backup and remove
	// file will always exist in case of updates, check to support generic usage
	if fileExists {
		if showLogs {
			fmt.Printf("> Creating backup of existing file (%s)\n", source)
		}
		backupTargetFile := filepath.Dir(source) + filepath.Base(source) + "-backup"
		if err = CopyFile(target, backupTargetFile); err != nil {
			return err
		}

		if err = os.Remove(target); err != nil {
			return err
		}

		defer func() {
			// remove backup: when no panic
			if panicErr := recover(); panicErr != nil {
				err = fmt.Errorf("failed to move file, failed to restore backup: %v", panicErr)
			} else {
				if showLogs {
					fmt.Println("> Removing backup file")
				}
				os.Remove(backupTargetFile)
			}
		}()
	}

	err = os.Rename(source, target)
	if err != nil {
		// if file existed initially, place it back
		if fileExists {
			if showLogs {
				fmt.Println("> Failed to move updated file, restoring from backup")
			}

			backupTargetFile := filepath.Dir(source) + filepath.Base(source) + "-backup"
			if backupErr := os.Rename(backupTargetFile, target); backupErr != nil {
				fmt.Println()
				fmt.Printf("Unable to restore original file \nKindly move %s to %s \n", backupTargetFile, target)
				fmt.Println()
				// panic to skip deletion of
				panic(backupTargetFile)
			}
		}
		return err
	}
	if showLogs {
		fmt.Println("> Move successful")
	}

	return nil
}
