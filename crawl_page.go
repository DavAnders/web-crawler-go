package main

import (
	"fmt"
	"net/url"
	"strings"
)

func (cfg *config) crawlPage(rawCurrentURL string) error {
	// If the maximum number of pages has been reached, return nil to stop the recursion
	if len(cfg.pages) >= cfg.maxPages {
		cfg.wg.Done()
		return nil
	}

	defer cfg.wg.Done() // Decrement the WaitGroup counter when the function completes in order to signal to the main goroutine that it has finished

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		return fmt.Errorf("unable to parse current URL: %v", err)
	}

	if cfg.baseURL.Hostname() != currentURL.Hostname() {
		fmt.Printf("Skipping %v\n", rawCurrentURL)
		return nil
	}

	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		return fmt.Errorf("unable to normalize URL: %v", err)
	}

	if isFirstVisit := cfg.addPageVisit(normalizedURL); !isFirstVisit {
		return nil
	}

	// Get the HTML content of the page with print statement
	htmlBody, err := getHTML(normalizedURL)
	if err != nil {
		// Skip the URL if the content type is not HTML
		if strings.Contains(err.Error(), "unexpected content type") {
			fmt.Printf("Skipping %v due to content type\n", normalizedURL)
			return nil
		}
		return fmt.Errorf("unable to get HTML: %v", err)
	}
	fmt.Printf("Crawling %v\n", normalizedURL)
	fmt.Println(htmlBody)

	// Get all the URLs from the HTML content
	urls, err := getURLsFromHTML(htmlBody, normalizedURL)
	if err != nil {
		return fmt.Errorf("unable to get URLs from HTML: %v", err)
	}

	// Recursively crawl the URLs
	for _, url := range urls {
		cfg.wg.Add(1) // Increment the WaitGroup counter to signal a new goroutine
		go func(url string) {
			cfg.concurrencyControl <- struct{}{}        // Limit the number of concurrent goroutines
			defer func() { <-cfg.concurrencyControl }() // Release the semaphore when the goroutine completes
			cfg.crawlPage(url)
		}(url)
	}
	return nil
}

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	if _, exists := cfg.pages[normalizedURL]; exists {
		cfg.pages[normalizedURL]++
		return false
	}
	cfg.pages[normalizedURL] = 1
	return true
}
