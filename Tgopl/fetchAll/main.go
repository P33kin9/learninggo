package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"
)

func fetchUrl(url string, ch chan<- map[string]string) {
	start := time.Now()
	resp, err := http.Get(url)
	result := make(map[string]string)
	if err != nil {
		result[url] = fmt.Sprintf("http-get: %v", err)
		ch <- result
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		result[url] = fmt.Sprintf("while reading:%v %v", url, err)
		ch <- result
		return
	}

	resp.Body.Close()
	secs := time.Since(start).Seconds()
	result[url] = fmt.Sprintf("%v %v %v Bytes %v's", url, resp.Status, nbytes, secs)
	ch <- result
}

func main() {
	start := time.Now()
	ch := make(chan map[string]string)
	count := 0
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "https://") {
			url = "https://" + url
		}
		fmt.Printf("start fetch %v\n", url)
		go fetchUrl(url, ch)
		go fetchUrl(url, ch)
		go fetchUrl(url, ch)
		count += 3
	}

	result := make(map[string]string)
	keys := []string{}
	for i := 0; i < count; i++ {
		for k, v := range <-ch {
			result[k] = v
			keys = append(keys, k)
		}
	}

	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("%s: %s\n", k, result[k])
	}
	fmt.Printf("done, use %v seconds", time.Since(start).Seconds())
}
