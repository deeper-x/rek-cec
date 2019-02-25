package checker

import (
	"log"
	"net/http"
)

// CallService check if a given url is available
// Return a boolean
func CallService(url string, expectedStatus int) (bool, string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Fatal error calling service...", err)
	}

	// Status code acceptable
	if resp.StatusCode == expectedStatus {
		return true, url
	}

	return false, url
}
