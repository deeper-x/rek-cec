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
func WriteResponseOutput(exitStatus bool, url string, expectedStatus int) string {
	response := "# ERROR #"

	if exitStatus {
		response = "WORKS!"
	}

	return fmt.Sprintf("%v [code %v]: %v", url, expectedStatus, response)
}

// LoadConfigurationFile read configuration data
func LoadConfigurationFile(inFile string) (Config, error) {
	var config Config
	configFile, err := os.Open(inFile)

	defer configFile.Close()

	if err != nil {
		fmt.Println(err.Error())
		return config, err
	}

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)

	return config, nil
}

// GetExpectedStatus check input parameter
func GetExpectedStatus() (int, error) {
	file := fmt.Sprintf("%v/src/github.com/%v/rek-cec/settings/configuration.json", os.Getenv("GOPATH"), os.Getenv("USER"))
	configFile, err := LoadConfigurationFile(file)

	if err != nil {
		return -1, err
	}
	expectedStatus := configFile.ServiceAvailable

	argsLen := len(os.Args)

	if argsLen == 2 {
		var err error
		expectedStatus, err = strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatal("Error parameter: ", err)
		}
	}

	return expectedStatus, nil
}
