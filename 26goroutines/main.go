package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup // usually they are pointers

func main() {
	websitelist := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
	}
	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait()
}

func getStatusCode(endpoint string) {

	defer wg.Done()

	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("OOPS there is an error")
	}
	fmt.Printf("%d status code for %s \t", res.StatusCode, endpoint)
}
