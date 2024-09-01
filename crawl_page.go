package main

import (
	"fmt"
	"net/url"
	"strings"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) (map[string]int, error) {
	// Check if current URL is on the same domain as the base URL
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return map[string]int{}, fmt.Errorf("unable to parse base URL: %v", err)
	}

	// Parse the current URL to compare hostnames
	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		return map[string]int{}, fmt.Errorf("unable to parse current URL: %v", err)
	}

	// If the hostnames are different, return the map as is
	if baseURL.Hostname() != currentURL.Hostname() {
		fmt.Printf("Skipping %v\n", rawCurrentURL)
		return pages, nil
	}

	// Normalize the URL for comparison
	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		return map[string]int{}, fmt.Errorf("unable to normalize URL: %v", err)
	}

	// Check if the URL is already in the map
	if _, exists := pages[normalizedURL]; exists {
		pages[normalizedURL]++
		return pages, nil
	}
	pages[normalizedURL] = 1

	// Get the HTML content of the page with print statement
	htmlBody, err := getHTML(normalizedURL)
	if err != nil {
		// Skip the URL if the content type is not HTML
		if strings.Contains(err.Error(), "unexpected content type") {
			fmt.Printf("Skipping %v\n", normalizedURL)
			return pages, nil
		}
		return map[string]int{}, fmt.Errorf("unable to get HTML: %v", err)
	}
	fmt.Printf("Crawling %v\n", normalizedURL)
	fmt.Println(htmlBody)

	// Get all the URLs from the HTML content
	urls, err := getURLsFromHTML(htmlBody, normalizedURL)
	if err != nil {
		return map[string]int{}, fmt.Errorf("unable to get URLs from HTML: %v", err)
	}

	// Recursively crawl the URLs
	for _, url := range urls {
		fmt.Printf("Found URL: %v\n", url)
		fmt.Println("Crawling...")
		pages, err = crawlPage(rawBaseURL, url, pages)
		if err != nil {
			return map[string]int{}, fmt.Errorf("unable to crawl page: %v", err)
		}
	}

	return pages, nil
}
