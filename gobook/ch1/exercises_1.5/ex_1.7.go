// fetch URL like cURL

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const output_filename = "fetched_content.out"

func main() {
	for _, url := range os.Args[1:] {
		response, err := http.Get(url)
		if err != nil {
			fmt.Fprint(os.Stderr, "fetching: %v \n", err)
			os.Exit(1)
		}

		output, err := os.Create(output_filename)
		if err != nil {
			fmt.Fprint(os.Stderr, "fetching: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		// copying results to file
		copyLen, err := io.Copy(output, response.Body)
		if err != nil {
			fmt.Fprint(os.Stderr, "copying: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		// closing resources
		output.Close()
		response.Body.Close()
		fmt.Printf("saved %d bytes into %s file!", copyLen, output_filename)
	}
}
