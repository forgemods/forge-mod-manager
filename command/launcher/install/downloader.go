package install

import (
	"io"
	"net/http"
	"os"
)

func DownloadFile(filepath string, url string) error {
	if _, err := os.Stat(filepath); err == nil {
		//log.Printf("file exists, skipping. `%s`", filepath)
		return nil
	}

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer Close(resp.Body)

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}

	// Write the body to file
	_, err = io.Copy(out, resp.Body)

	if err != nil {
		return err
	}

	Close(out)

	return nil
}