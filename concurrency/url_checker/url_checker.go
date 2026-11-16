// Program to check several websites at the same time using goroutines
// and a channel to collect the results.
package main

import (
	"fmt"
	"net/http"
	"time"
)

// A client with a timeout so one slow site cannot block the program forever
var client = &http.Client{Timeout: 10 * time.Second}

// checkURL sends the status of url to ch
func checkURL(url string, ch chan<- string) {
	resp, err := client.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("Site %s is down. Error: %v", url, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		ch <- fmt.Sprintf("Site %s is up. Status code: %d", url, resp.StatusCode)
	} else {
		ch <- fmt.Sprintf("Site %s is down. Status code: %d", url, resp.StatusCode)
	}
}

func main() {
	urls := []string{
		"https://example.com",
		"https://example.org",
		"https://example.net",
		"https://go.dev",
	}

	results := make(chan string)

	// Check every URL in its own goroutine
	for _, url := range urls {
		go checkURL(url, results)
	}

	// Receive exactly one result per URL
	for range urls {
		fmt.Println(<-results)
	}
}
