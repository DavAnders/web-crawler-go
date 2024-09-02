package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"strconv"
)

func main() {
	baseURL, maxConcurrency, maxPages, err := parseArguments(os.Args)
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

	// Create a new config with a concurrency limit and maximum number of pages
	cfg := newConfig(parsedBaseURL, maxConcurrency, maxPages)

	// Start the crawl with the base URL
	cfg.wg.Add(1)
	go cfg.crawlPage(baseURL)

	// Wait for all the crawlers to finish
	cfg.wg.Wait()

	printCrawlSummary(cfg.pages)
}

// parseArguments validates the arguments provided to the program
func parseArguments(args []string) (string, int, int, error) {
	var maxConcurrency, maxPages int

	// Set default values
	maxConcurrency = 5
	maxPages = 100

	// Define the flag set
	flagSet := flag.NewFlagSet(args[0], flag.ContinueOnError)
	flagSet.IntVar(&maxConcurrency, "concurrency", maxConcurrency, "maximum number of concurrent goroutines")
	flagSet.IntVar(&maxPages, "maxpages", maxPages, "maximum number of pages to crawl")

	// Parse the flags from the command-line arguments
	err := flagSet.Parse(args[1:])
	if err != nil {
		return "", 0, 0, fmt.Errorf("unable to parse flags: %v", err)
	}

	// Get remaining non-flag arguments
	remainingArgs := flagSet.Args()
	if len(remainingArgs) == 0 {
		return "", 0, 0, fmt.Errorf("no website provided")
	}

	// Extract base URL and optionally override maxConcurrency and maxPages
	baseURL := remainingArgs[0]

	if len(remainingArgs) > 1 {
		if concurrency, err := strconv.Atoi(remainingArgs[1]); err == nil {
			maxConcurrency = concurrency
		} else {
			return "", 0, 0, fmt.Errorf("invalid concurrency value: %v", remainingArgs[1])
		}
	}

	if len(remainingArgs) > 2 {
		if pages, err := strconv.Atoi(remainingArgs[2]); err == nil {
			maxPages = pages
		} else {
			return "", 0, 0, fmt.Errorf("invalid maxpages value: %v", remainingArgs[2])
		}
	}

	return baseURL, maxConcurrency, maxPages, nil
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
