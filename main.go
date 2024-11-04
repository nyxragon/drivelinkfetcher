package main

import (
	"drivelinkfetcher"
	"log"
)

func main() {
	apiKey := "your_api_key"
	cx := "your_cx_key"
	query := "file"
	from := "2024-01-01-00"
	to := "2024-12-31-23"
	includeLabel := true
	outputFile := "output.txt"

	if err := drivelinkfetcher.FetchDriveLinks(apiKey, cx, query, from, to, includeLabel, outputFile); err != nil {
		log.Fatalf("Error: %v", err)
	}
}
