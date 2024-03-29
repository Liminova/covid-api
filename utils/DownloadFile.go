package utils

import (
	"io"
	"net/http"
	"os"
)

func DownloadFile(filepath string, url string) error {
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Ovewrite the file
	if _, err := io.Copy(out, resp.Body); err != nil {
		return err
	}

	return nil
}
