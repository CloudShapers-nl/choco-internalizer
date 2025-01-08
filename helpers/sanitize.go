package helpers

import (
	"fmt"
	"os"
)

func SanitizePackageDir(pkgsDir string) error {
	entries, err := os.ReadDir(pkgsDir)
	if err != nil {
		return err
	}

	for _, e := range entries {
		filePath := fmt.Sprintf("%s/%s", pkgsDir, e.Name())

		if e.Name() == "_rels" {
			err := os.RemoveAll(filePath)
			if err != nil {
				return err
			}
			continue
		}

		if e.Name() == "package" {
			err := os.RemoveAll(filePath)
			if err != nil {
				return err
			}
			continue
		}

		if e.Name() == "[Content_Types].xml" {
			err := os.Remove(filePath)
			if err != nil {
				return err
			}
			continue
		}

	}

	return nil
}
