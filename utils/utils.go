package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Config data
type Config struct {
	ServiceListFile  string
	ServiceAvailable int
}

// WriteResponseOutput write a response string
func WriteResponseOutput(exitStatus bool, url string, expectedStatus int) {
	response := "# ERROR #"

	if exitStatus {
		response = "WORKS!"
	}

	fmt.Printf("%v [code %v]: %v\n", url, expectedStatus, response)
}

// LoadConfigurationFile read configuration data
func LoadConfigurationFile(inFile string) Config {
	var config Config
	configFile, err := os.Open(inFile)

	defer configFile.Close()

	if err != nil {
		fmt.Println(err.Error())
	}

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)

	return config
}

// GetExpectedStatus check input parameter
func GetExpectedStatus() int {
	configFile := LoadConfigurationFile("./settings/configuration.json")
	expectedStatus := configFile.ServiceAvailable

	argsLen := len(os.Args)

	if argsLen == 2 {
		var err error
		expectedStatus, err = strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatal("Error parameter: ", err)
		}
	}

	return expectedStatus
}
