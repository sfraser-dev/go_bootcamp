package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	var websites = []string{
		"https://google.com",
		"https://www.facebook.com",
		"https://go.dev",
		"https://stackoverflow.com",
		"https://www.amazon.co.uk",
	}

	// make creates a value out of the given type
	c := make(chan string)

	for _, website := range websites {
		//checkLink(website)
		go checkLink(website, c)
	}

	/* recieve all the channel messages(iterate over the channel) */
	//for { go checkLink(<-c, c) }
	/* refactor: wait for the channel to return some value, assign to ws */
	//for ws := range c { go checkLink(ws, c) }
	/* refactor to use function literal (lambda / anonymous function) -
	function literals are handy in go routine invocations (as Go Routines
	must operate on a function) */
	for ws := range c {
		go func(wsFuncLit string) {
			time.Sleep(2 * time.Second)
			checkLink(wsFuncLit, c)
		}(ws) /* note the parentheses - must call the function */
	}
}

func checkLink(website string, c chan string) {
	_, err := http.Get(website)
	if err != nil {
		msg := fmt.Sprintf("%s: %s\n", website, err)
		fmt.Print(msg)
		c <- website
		return
	}
	// io.Copy(os.Stdout, resp.Body)
	msg := fmt.Sprintf("%s: yip, it's up\n", website)
	fmt.Print(msg)
	c <- website
}
