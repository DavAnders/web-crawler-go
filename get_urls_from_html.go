package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return []string{}, fmt.Errorf("unable to parse base URL: %v", err)
	}

	reader := strings.NewReader(htmlBody)

	tree, err := html.Parse(reader)
	if err != nil {
		return []string{}, fmt.Errorf("unable to parse HTML: %v", err)
	}

	// Initialize the slice as empty
	urls := []string{}

	// Traverse the HTML tree and extract URLs
	var traverseNodeTree func(node *html.Node)
	traverseNodeTree = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "a" {
			// Extract href attribute from the node
			for _, attr := range node.Attr {
				if attr.Key == "href" {
					parsedURL, err := url.Parse(attr.Val)
					if err != nil {
						fmt.Printf("unable to parse href '%v': %v\n", attr.Val, err)
						continue
					}

					resolvedURL := baseURL.ResolveReference(parsedURL)
					urls = append(urls, resolvedURL.String())
				}
			}
		}
		// Recursively traverse the tree, depth-first
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			traverseNodeTree(child)
		}
	}
	traverseNodeTree(tree)

	return urls, nil
}
