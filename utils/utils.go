package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

// Config data
type Config struct {
	ServiceListFile string
}

// WriteResponseOutput write a response string
func WriteResponseOutput(exitStatus bool, url string) {
	if exitStatus {
		fmt.Println(url, "success")
	} else {
		fmt.Println(url, "failure")
	}
}

// LoadConfigurationData read configuration data
func LoadConfigurationData(inFile string) Config {
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
