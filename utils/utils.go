package utils

import (
	"encoding/json"
	"fmt"
	"os"
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

	fmt.Printf("%v [%v]: %v\n", url, expectedStatus, response)
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
