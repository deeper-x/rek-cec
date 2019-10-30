package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/deeper-x/rek-cec/checker"
	"github.com/deeper-x/rek-cec/utils"
)

func main() {
	expectedStatus, err := utils.GetExpectedStatus()
	if err != nil {
		log.Println(err)
	}

	configFile, err := utils.LoadConfigurationFile("./settings/configuration.json")
	if err != nil {
		log.Println(err)
	}

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

	for scanner.Scan() {
		isAvailable, url := checker.CallService(scanner.Text(), expectedStatus)
		fmt.Println(utils.WriteResponseOutput(isAvailable, url, expectedStatus))
	}
}
