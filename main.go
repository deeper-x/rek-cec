package main

import (
	"bufio"
	"log"
	"os"

	"github.com/deeper-x/rek-cec/checker"
	"github.com/deeper-x/rek-cec/utils"
)

func main() {
	configFile := utils.LoadConfigurationFile("./assets/configuration.json")

	urlFile := configFile.ServiceListFile
	fileIn, err := os.Open(urlFile)
	if err != nil {
		log.Fatal("Error opening file!", err)
	}

	defer fileIn.Close()

	scanner := bufio.NewScanner(fileIn)

	if err := scanner.Err(); err != nil {
		log.Fatal("Error fetching file content!", err)
	}

	expectedStatus := configFile.ServiceAvailable
	for scanner.Scan() {
		isAvailable, url := checker.CallService(scanner.Text(), expectedStatus)
		utils.WriteResponseOutput(isAvailable, url, expectedStatus)
	}
}
