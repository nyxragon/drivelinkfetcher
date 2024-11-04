package drivelinkfetcher

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"
)

const apiURL = "https://www.googleapis.com/customsearch/v1"

// SearchResult holds the structure of a single search result
type SearchResult struct {
	Title string `json:"title"`
	Link  string `json:"link"`
}

// SearchResponse holds the structure of the API response
type SearchResponse struct {
	Items []SearchResult `json:"items"`
}

// FetchDriveLinks performs a Google Custom Search with the specified parameters and returns formatted results.
func FetchDriveLinks(apiKey, cx, query, from, to string, includeLabel bool, outputFile string) error {
	dateRange, err := CalculateDateRange(from, to)
	if err != nil {
		return err
	}

	results, err := SearchGoogle(apiKey, cx, query, dateRange)
	if err != nil {
		return err
	}

	output := FormatResults(results, includeLabel)

	if outputFile != "" {
		return os.WriteFile(outputFile, []byte(output), 0644)
	}
	fmt.Println(output)
	return nil
}

// SearchGoogle queries the Google Custom Search API
func SearchGoogle(apiKey, cx, query, dateRange string) (*SearchResponse, error) {
	reqURL := fmt.Sprintf("%s?key=%s&cx=%s&q=%s&dateRestrict=%s", apiURL, apiKey, cx, url.QueryEscape(query), dateRange)

	resp, err := http.Get(reqURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %s", resp.Status)
	}

	var searchResponse SearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&searchResponse); err != nil {
		return nil, err
	}

	return &searchResponse, nil
}

// CalculateDateRange parses a date range in the format yyyy-mm-dd-hh
func CalculateDateRange(from, to string) (string, error) {
	const layout = "2006-01-02-15"
	var startDate, endDate time.Time
	var err error

	if from == "" && to == "" {
		endDate = time.Now()
		startDate = endDate.AddDate(0, -6, 0)
	} else {
		startDate, err = time.Parse(layout, from)
		if err != nil {
			return "", fmt.Errorf("invalid 'from' date format, expected yyyy-mm-dd-hh")
		}
		endDate, err = time.Parse(layout, to)
		if err != nil {
			return "", fmt.Errorf("invalid 'to' date format, expected yyyy-mm-dd-hh")
		}
	}

	return fmt.Sprintf("%s,%s", startDate.Format("2006-01-02-15"), endDate.Format("2006-01-02-15")), nil
}

// FormatResults formats results based on includeLabel flag
func FormatResults(results *SearchResponse, includeLabel bool) string {
	var output string
	for _, item := range results.Items {
		if includeLabel {
			output += fmt.Sprintf("Title: %s\nURL: %s\n\n", item.Title, item.Link)
		} else {
			output += fmt.Sprintf("%s\n", item.Link)
		}
	}
	return output
}
