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

	path := strings.TrimSuffix(parsedURL.Path, "/")

	normalizedURL := parsedURL.Hostname() + path

	return normalizedURL, nil
}
