package checker

import (
	"log"
	"net/http"
)

// CallService check if a given url is available
// Return a boolean
func CallService(url string) (bool, string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Fatal error calling http service...", err)
	}

	// Status code acceptable is the 200 OK
	if resp.StatusCode == 200 {
		return true, url
	}

	return false, url
}
