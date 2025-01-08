package helpers

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func DownloadFile(name string, version string) (string, error) {
	filePath := fmt.Sprintf("./%s_%s.nupkg", name, version)

	out, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer out.Close()

	fileUrl := fmt.Sprintf("https://community.chocolatey.org/api/v2/package/%s/%s", name, version)

	resp, err := http.Get(fileUrl)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("NON OK status code: %d", resp.StatusCode)
	}

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", err
	}

	return filePath, nil
}
