package utils

import "fmt"

// WriteResponseOutput write a response string
func WriteResponseOutput(exitStatus bool, url string) {
	if exitStatus {
		fmt.Println(url, "success")
	} else {
		fmt.Println(url, "failure")
	}
}
