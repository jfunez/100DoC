// fetch URL like cURL

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		response, err := http.Get(url)
		if err != nil {
			fmt.Fprint(os.Stderr, "fetching: %v \n", err)
			os.Exit(1)
		}
		fmt.Printf("status code: %s\n", response.Status)
		body, err := ioutil.ReadAll(response.Body)
		response.Body.Close()
		if err != nil {
			fmt.Fprint(os.Stderr, "fetching: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", body)
	}
}
