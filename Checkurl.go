package main

import (
	"fmt"
	"net/http"
)

// Function to check if a URL is up
func checkURL(url string, ch chan<- string) {
	// Make a GET request
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("Site %s is down. Error: %v", url, err)
		return
	}
	defer resp.Body.Close()

	// Check the status code
	if resp.StatusCode == http.StatusOK {
		ch <- fmt.Sprintf("Site %s is up. Status code: %d", url, resp.StatusCode)
	} else {
		ch <- fmt.Sprintf("Site %s is down. Status code: %d", url, resp.StatusCode)
	}
}

func main() {
	// List of URLs to check
	urls := []string{
		"http://example.com",
		"http://example.org",
		"http://example.net",
		"http://example.edu",
	}

	// Channel to receive results
	results := make(chan string)

	// Check each URL
	for _, url := range urls {
		go checkURL(url, results)
	}

	// Print results
	for range urls {
		fmt.Println(<-results)
	}
}
