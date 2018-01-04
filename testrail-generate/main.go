package main

import (
	"os"

	"github.com/educlos/testrail"
)

func main() {
	url := os.Getenv("TESTRAIL_URL")
	if url == "" {
		url = "http://localhost:7070/testrail"
	}
	username := os.Getenv("TESTRAIL_USERNAME")
	password := os.Getenv("TESTRAIL_TOKEN")

	client := testrail.NewClient(url, username, password)

	err := client.GenerateCustom()
	if err != nil {
		panic(err)
	}

}
