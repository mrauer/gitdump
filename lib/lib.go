package lib

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func DownloadFile(filepath string, url string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func MakeDir(entity string) (string, error) {
	path := fmt.Sprintf("data/%s/%s", entity, time.Now().Format("2006-01-02"))
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, 0700)
		return path, nil
	}
	return "", nil
}
