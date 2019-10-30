package utils

import (
	"fmt"
	"os"
	"testing"
)

func TestWriteResponseOutput(t *testing.T) {
	res := WriteResponseOutput(true, "http://yahoo.com", 200)

	if len(res) == 0 {
		t.Error("response output")
	}
}

func TestLoadConfigurationFile(t *testing.T) {
	file := fmt.Sprintf("%v/src/github.com/%v/rek-cec/settings/configuration.json", os.Getenv("GOPATH"), os.Getenv("USER"))
	_, err := LoadConfigurationFile(file)

	if err != nil {
		t.Error("Configuration file not loaded")
	}

}

func TestGetExpectedStatus(t *testing.T) {
	_, err := GetExpectedStatus()
	if err != nil {
		t.Error("Expected status not read")
	}

}
