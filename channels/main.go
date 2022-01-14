package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	urls := []string{"https://google.com", "https://facebook.com", "https://golang.org", "https://amazon.com", "https://stackoverflow.com"}

	c := make(chan string)

	for _, url := range urls {
		go checkUrl(url, c)
	}
	for l := range c {
		go func(url string) {
			time.Sleep(5 * time.Second)
			checkUrl(url, c)
		}(l)
	}
}

func checkUrl(url string, c chan string) {
	fmt.Println("Requesting for", url)
	_, err := http.Get(url)

	if err != nil {
		fmt.Println("Failed to fetch", url, "Error:", err)
		c <- url
		return
	}

	fmt.Println(url, "is up!")
	c <- url
}
