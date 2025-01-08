package helpers

import (
	"os/exec"
)

func PackAndPush(pkgsDir string, chocoRepo string, apiKey string) error {
	chocoPack := exec.Command("choco", "pack")
	chocoPack.Dir = pkgsDir
	_, err := chocoPack.Output()
	if err != nil {
		return err
	}

	chocoPush := exec.Command("choco", "push", "--source", chocoRepo, "--api-key", apiKey)
	chocoPush.Dir = pkgsDir
	_, err = chocoPush.Output()
	if err != nil {
		return err
	}

	return nil
}
