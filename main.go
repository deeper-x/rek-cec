package main

import (
	"bufio"
	"log"
	"os"

	"github.com/deeper-x/wserver/checker"
	"github.com/deeper-x/wserver/utils"
)

func main() {
	fileIn, err := os.Open("./assets/to_check.txt")
	if err != nil {
		log.Fatal("Error opening file!", err)
	}

	defer fileIn.Close()

	scanner := bufio.NewScanner(fileIn)

	if err := scanner.Err(); err != nil {
		log.Fatal("Error fetching file content!", err)
	}

	for scanner.Scan() {
		isAvailable, url := checker.CallService(scanner.Text())
		utils.WriteResponseOutput(isAvailable, url)
	}

}
