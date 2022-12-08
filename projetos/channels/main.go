package main

import (
	"fmt"
	"net/http"
)

func main() {

	links := []string{
		"http://www.google.com",
		"http://www.facebook.com",
		"http://www.stackoverflow.com",
		"http://www.golang.org",
		"http://www.amazon.com",
	}

	for _, link := range links {
		checkLink(link)
	}
}

func checkLink(url string) {
	_, err := http.Get(url)

	if err != nil {
		fmt.Println(url, "===> might be down!")
	}

	fmt.Println(url, "===> is up!")
}
