# Drivelinkfetcher

Drivelinkfetcher is a Go package for performing Google Custom Searches and fetching relevant links. This package allows you to specify date ranges, customize output formats, and save results to a file.

## Installation

To use the `drivelinkfetcher` package, install it in your Go project:

```
go get github.com/yourusername/drivelinkfetcher
```

## Usage

### Importing the Package

In your Go program, import the package as follows:

```
import (
    "drivelinkfetcher"
    "log"
)
```

### FetchDriveLinks Function

Use the `FetchDriveLinks` function to perform a search. It requires the following parameters:

- **apiKey**: Your Google API key.
- **cx**: Your Google Custom Search engine ID.
- **query**: The search query (e.g., "file").
- **from**: Start date in the format `yyyy-mm-dd-hh` (e.g., "2024-01-01-00").
- **to**: End date in the format `yyyy-mm-dd-hh` (e.g., "2024-12-31-23").
- **includeLabel**: A boolean flag (true or false) to indicate whether to include titles in the output.
- **outputFile**: Optional. The name of the file to save results.

### Example Scenarios

#### Example 1: Basic Usage without Include Label

Fetch links without titles and save the output to `output.txt`:

```
func main() {
    apiKey := "your_api_key"
    cx := "your_cx_key"
    query := "file"
    from := "" // Defaults to last 6 months
    to := ""   // Defaults to now
    includeLabel := false
    outputFile := "output.txt"

    if err := drivelinkfetcher.FetchDriveLinks(apiKey, cx, query, from, to, includeLabel, outputFile); err != nil {
        log.Fatalf("Error: %v", err)
    }
}
```

#### Example 2: Usage with Include Label

Fetch links and include titles in the output:

```
func main() {
    apiKey := "your_api_key"
    cx := "your_cx_key"
    query := "file"
    from := "" // Defaults to last 6 months
    to := ""   // Defaults to now
    includeLabel := true
    outputFile := "output.txt"

    if err := drivelinkfetcher.FetchDriveLinks(apiKey, cx, query, from, to, includeLabel, outputFile); err != nil {
        log.Fatalf("Error: %v", err)
    }
}
```

#### Example 3: Custom Date Range

Specify a custom date range for the search:

```
func main() {
    apiKey := "your_api_key"
    cx := "your_cx_key"
    query := "file"
    from := "2024-01-01-00" // Start date
    to := "2024-12-31-23"   // End date
    includeLabel := true
    outputFile := "output.txt"

    if err := drivelinkfetcher.FetchDriveLinks(apiKey, cx, query, from, to, includeLabel, outputFile); err != nil {
        log.Fatalf("Error: %v", err)
    }
}
```

## License

This project is licensed under the MIT License. See the LICENSE file for details.

## Contributing

Contributions are welcome! Feel free to submit issues and pull requests.
