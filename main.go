package main

import (
	"fmt"
	"net/url"
	"os"
)

func main() {
	baseURL, err := parseArguments(os.Args)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Parse to make sure the URL is valid
	parsedBaseURL, err := url.Parse(baseURL)
	if err != nil {
		fmt.Println("Error parsing base URL:", err)
		os.Exit(1)
	}

	// Create a new config with a concurrency limit of 5
	cfg := newConfig(parsedBaseURL, 5)

	// Start the crawl with the base URL
	cfg.wg.Add(1)
	go cfg.crawlPage(baseURL)

	// Wait for all the crawlers to finish
	cfg.wg.Wait()

	printCrawlSummary(cfg.pages)
}

// parseArguments validates the arguments provided to the program
func parseArguments(args []string) (string, error) {
	if len(args) < 2 {
		return "", fmt.Errorf("no website provided")
	}

	if len(args) > 2 {
		return "", fmt.Errorf("too many arguments provided")
	}

	return args[1], nil
}

// printCrawlSummary prints the results of the crawl
func printCrawlSummary(pages map[string]int) {
	fmt.Println("\nCrawl Summary:")
	for page, count := range pages {
		if count == 1 {
			fmt.Printf("%v: crawled %v time\n", page, count)
		} else {
			fmt.Printf("%v: crawled %v times\n", page, count)
		}
	}
	fmt.Printf("Total pages crawled: %v\n", len(pages))
}
