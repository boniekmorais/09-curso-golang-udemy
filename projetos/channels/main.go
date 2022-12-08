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

	// Creating a channel (Type string)
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	fmt.Println(<-c)
}

func checkLink(url string, c chan string) {

	_, err := http.Get(url)

	if err != nil {
		fmt.Println(url, "===> might be down!")
		c <- "Might be down I think"
	}

	fmt.Println(url, "===> is up!")

	c <- "Yes its up!"

	// time.Sleep(10 * time.Second)
}
