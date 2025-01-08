package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/cloudshapers-nl/choco-internalizer/helpers"
)

func main() {
	list, err := helpers.GetPackageList()
	if err != nil {
		log.Fatal(err)
	}

	inputFile := os.Args[1]
	chocoRepo := os.Args[2]
	apiKey := os.Args[3]

	inputCont, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	var input map[string]string
	err = json.Unmarshal(inputCont, &input)
	if err != nil {
		log.Fatal(err)
	}

	for name, version := range input {
		exists := helpers.CheckIfExists(list, name, version)
		if exists {
			fmt.Printf("Package %s with version %s already exists in the repo\n", name, version)
			continue
		}

		fmt.Printf("Downloading package: %s, version: %s\n", name, version)

		filePath, err := helpers.DownloadFile(name, version)
		if err != nil {
			log.Fatal(err)
		}

		pkgsDir, err := helpers.Unzip(filePath)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Sanitzing package %s version %s\n", name, version)

		if err := helpers.SanitizePackageDir(pkgsDir); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Pushing package %s version %s\n", name, version)

		if err := helpers.PackAndPush(pkgsDir, chocoRepo, apiKey); err != nil {
			log.Fatal(err)
		}
	}
}
