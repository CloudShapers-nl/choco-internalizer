package helpers

import (
	"os/exec"
	"strings"
)

func GetPackageList() ([]map[string]string, error) {
	outputList := []map[string]string{}

	cmd := exec.Command("choco", "find", "--all-versions", "-r")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(output), "\r\n")

	for _, line := range lines {
		if line == "" {
			continue
		}
		slices := strings.Split(line, "|")

		outputMap := map[string]string{
			strings.ToLower(slices[0]): slices[1],
		}

		outputList = append(outputList, outputMap)
	}

	return outputList, nil
}

func CheckIfExists(list []map[string]string, packageName string, packageVersion string) bool {
	for _, pkgs := range list {
		for k, v := range pkgs {
			if k == packageName && v == packageVersion {
				return true
			}
		}
	}
	return false
}
