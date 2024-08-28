package main

import "fmt"

func main() {
	str, err := normalizeURL("https://blog.boot.dev/path/")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str)
}
