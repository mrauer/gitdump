package lib

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/google/go-github/github"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

const (
	permissionMode = 0700
	dateFormat     = "2006-01-02"
	perPage        = 100
)

// DownloadFile downloads a file from a given URL and saves it to the specified filepath.
func DownloadFile(filePath string, url string) error {
	// Create the directory
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, permissionMode); err != nil {
		return fmt.Errorf("failed to create directory: %v", err)
	}

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to get data from URL: %v", err)
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to copy data to file: %v", err)
	}

	return nil
}

// MakeDir creates a directory with the specified entity name and today's date as the subdirectory.
func MakeDir(entity string) (string, error) {
	downloadPath := fmt.Sprintf("%s", viper.Get("download_path"))
	path := filepath.Join(downloadPath, entity, time.Now().Format(dateFormat))
	if err := os.MkdirAll(path, permissionMode); err != nil {
		return "", fmt.Errorf("failed to create directory: %v", err)
	}

	return path, nil
}

// GitLogin authenticates with GitHub using the provided token.
// It returns a context and a GitHub client.
func GitLogin() (context.Context, *github.Client) {
	token := fmt.Sprint(viper.Get("github_token"))
	if token == "" {
		log.Fatal("GitHub token is not provided. Set 'github_token' in your configuration.")
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	return ctx, client
}
