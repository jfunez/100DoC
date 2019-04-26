// fetch URL like cURL

package main

import (
	"strings"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {

		if !strings.HasPrefix(url, "http") {
			url = "http://" + url
		}
		fmt.Printf("url: %s", url)

		response, err := http.Get(url)
		if err != nil {
			fmt.Fprint(os.Stderr, "fetching: %v \n", err)
			os.Exit(1)
		}
		body, err := ioutil.ReadAll(response.Body)
		response.Body.Close()
		if err != nil {
			fmt.Fprint(os.Stderr, "fetching: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", body)
	}
}
