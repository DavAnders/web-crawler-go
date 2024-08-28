package main

import (
	"reflect"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
		<html>
			<body>
				<a href="/path/one">
					<span>Boot.dev</span>
				</a>
				<a href="https://other.com/path/one">
					<span>Boot.dev</span>
				</a>
			</body>
		</html>
		`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
		{
			name:     "all <a> tags are found",
			inputURL: "https://example.com",
			inputBody: `
			<html>
				<body>
					<a href="/path/one">
						<span>Link 1</span>
					</a>
					<a href="/path/two">
						<span>Link 2</span>
					</a>
					<a href="/path/three">
						<span>Link 3</span>
					</a>
				</body>
			</html>
			`,
			expected: []string{"https://example.com/path/one", "https://example.com/path/two", "https://example.com/path/three"},
		},
		{
			name:     "no <a> tags found, return empty slice",
			inputURL: "https://example.com",
			inputBody: `
			<html>
				<body>
					<p>No links here</p>
				</body>
			</html>
			`,
			expected: []string{},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getURLsFromHTML(tc.inputBody, tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}

			t.Logf("Test %d - '%s' actual: %v, expected: %v", i+1, tc.name, actual, tc.expected)
			if len(actual) != len(tc.expected) {
				t.Errorf("Test %v - %s FAIL: length mismatch, expected length: %d, actual length: %d", i, tc.name, len(tc.expected), len(actual))
			}

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			} else {
				t.Logf("Test %d - '%s' PASSED: expected URL: '%v', got: '%v'", i+1, tc.name, tc.expected, actual)
			}
		})
	}
}
