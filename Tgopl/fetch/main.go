package main

import (
	"fmt"
	"io"
	"strings"

	//"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "http-get: %v\n", err)
			os.Exit(10)
		}
		fmt.Printf("http-get: %v\n", resp.Status)
		//		b, err := ioutil.ReadAll(resp.Body)
		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(10)
		}
		resp.Body.Close()
		//fmt.Printf("%s", b)
	}
}
