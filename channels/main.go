package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
		"http://naver.com",
	}

	ch := make(chan string)

	for _, link := range links {
		go checkLink(link, ch)
		//fmt.Println(<-ch)
	}

	//for i := 0; i < len(links); i++ {
	//	fmt.Println(<-ch)
	//}
	// <- ch is a blocking operation
	// the main function hangs permanently for no more comes into ch
	// fmt.Println(<-ch)

	// infinitely execute get requests to the links
	/*
	for {
		str :=<- ch
		fmt.Println(str)
		link := strings.Split(str, " ")[0]
		go checkLink(link, ch)
	}
	 */

	// the equivalent code with line 33 ~ 38
	// Wait for channel to return some value.
	// After the channel returned some value,
	// assign it to l and run the body in for loop
	//for l := range ch {
	//	fmt.Println(l)
	//	go checkLink(strings.Split(l, " ")[0], ch)
	//}

	// function literal
	// Now we pass "copied value", not the "reference"
	// i.e. link == l
	for l := range ch {
		go func(link string) {
			time.Sleep(5 * time.Second)
			fmt.Println(link)
			checkLink(strings.Split(link, " ")[0], ch)
		}(l)
	}


}

func checkLink(link string, ch chan string) {
	_, err := http.Get(link)
	if err != nil {
		//fmt.Println(link, "might be down")
		ch <- link + " might be down!"
		return
	}

	//fmt.Println(link, "is up!")
	ch <- link + " is up!"
}