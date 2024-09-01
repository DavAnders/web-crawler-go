package main

import (
	"fmt"
	"os"
)

func main() {
	baseURL, err := parseArguments(os.Args)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println("Starting crawl of:", baseURL)
	pages, err := crawlPage(baseURL, baseURL, map[string]int{})
	if err != nil {
		fmt.Println("Error during crawling:", err)
		os.Exit(1)
	}

	printCrawlSummary(pages)
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
