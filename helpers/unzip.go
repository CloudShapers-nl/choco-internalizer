package helpers

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func Unzip(filePath string) (string, error) {
	dirString := strings.Split(filePath, "_")[0]
	newDir := fmt.Sprintf("%s_%d", dirString, time.Now().Unix())

	err := os.Mkdir(newDir, 0777)
	if err != nil {
		return "", err
	}

	cmd := exec.Command("tar", "-xf", filePath, "-C", newDir)
	_, err = cmd.Output()
	if err != nil {
		return "", err
	}

	return newDir, nil
}
