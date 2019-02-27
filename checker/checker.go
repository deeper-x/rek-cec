package checker

import (
	"log"
	"net/http"
)

// CallService check if a given url is available
// Return a boolean
func CallService(url string, expectedStatus int) (bool, string) {
	resp, err := http.Get(url)
	asExpected := false

	if err != nil {
		log.Fatal("Fatal error calling service...", err)
	}

	// Status code acceptable
	if resp.StatusCode == expectedStatus {
		asExpected = true
	}

	return asExpected, url
}
