package main

import "fmt"

func main() {
	str, err := normalizeURL("https://blog.boot.dev/path/")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str)

	body := `
	<html>
    	<body>
        	<a href="https://blog.boot.dev"><span>Go to Boot.dev</span></a>
    	</body>
	</html>
	`
	url := "https://blog.boot.dev"

	urlList, err := getURLsFromHTML(body, url)
	fmt.Println(urlList)
}
