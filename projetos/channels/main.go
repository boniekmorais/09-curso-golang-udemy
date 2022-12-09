package main

import (
	"fmt"
	"net/http"
	"time"
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

	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }

	// Infinity loop
	// for {
	// 	go checkLink(<-c, c)
	// }

	// for l := range c {
	// 	go checkLink(l, c)
	// }

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}

}

func checkLink(url string, c chan string) {

	// time.Sleep(5 * time.Second)

	_, err := http.Get(url)

	if err != nil {
		fmt.Println(url, "===> might be down!")
		c <- url
	}

	fmt.Println(url, "===> is up!")

	c <- url

	// time.Sleep(10 * time.Second)
}
