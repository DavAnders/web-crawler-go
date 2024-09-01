package main

import (
	"net/url"
	"strings"
)

func normalizeURL(rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	// Trim the trailing slash from the path
	path := strings.TrimSuffix(parsedURL.Path, "/")

	// Reconstruct with scheme, hostname, path
	normalizedURL := parsedURL.Scheme + "://" + parsedURL.Hostname() + path

	// Add the query if it exists
	if parsedURL.RawQuery != "" {
		normalizedURL += "?" + parsedURL.RawQuery
	}

	return normalizedURL, nil
}
