package main

import "fmt"

func printReport(pages map[string]int, baseURL string) {
	fmt.Printf(`
=============================
REPORT for %v
=============================
	`+"\n", baseURL)

	// Sort pages map into a slice of structs
	type page struct {
		URL   string
		Count int
	}
	var pagesSlice []page

	for URL, count := range pages {
		pagesSlice = append(pagesSlice, page{URL, count})
	}

	// Sort the slice by count
	for i := 0; i < len(pagesSlice); i++ {
		for j := i + 1; j < len(pagesSlice); j++ {
			if pagesSlice[i].Count < pagesSlice[j].Count {
				pagesSlice[i], pagesSlice[j] = pagesSlice[j], pagesSlice[i]
			}
		}
	}

	// Print the sorted slice
	for _, page := range pagesSlice {
		if page.Count == 1 {
			fmt.Printf("Found %v internal link to %v\n", page.Count, page.URL)
		} else {
			fmt.Printf("Found %v internal links to %v\n", page.Count, page.URL)
		}
	}

}
