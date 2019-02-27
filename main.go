package main

import (
	"bufio"
	"log"
	"os"
	"strconv"

	"github.com/deeper-x/rek-cec/checker"
	"github.com/deeper-x/rek-cec/utils"
)

func main() {
	configFile := utils.LoadConfigurationFile("./settings/configuration.json")
	expectedStatus := configFile.ServiceAvailable
	argsLen := len(os.Args)

	if argsLen == 2 {
		var err error
		expectedStatus, err = strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatal("Error parameter: ", err)
		}
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
		utils.WriteResponseOutput(isAvailable, url, expectedStatus)
	}
}
