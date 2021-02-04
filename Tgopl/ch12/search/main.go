// Search is a demo of the params.Unpack function.
package main

import (
	"cf/ch12/params" //!+
	"fmt"
	"log"
	"net/http"
)

// search implements the /search URL endpoint.
func search(resp http.ResponseWriter, req *http.Request) {
	var data struct {
		Labels     []string `http:"l"`
		MaxResults int      `http:"max"`
		Exact      bool     `http:"x"`
	}
	data.MaxResults = 10 // set default
	if err := params.Unpack(req, &data); err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest) // 400
		return
	}

	// ... rest of handler...
	fmt.Fprintf(resp, "Search: %+v\n", data)
}

//!-

func main() {
	http.HandleFunc("/search", search)
	log.Fatal(http.ListenAndServe(":12345", nil))
}

/*
//!+output
$ gor main.go 'http://localhost:12345/search'
http-get: 200 OK
Search: {Labels:[] MaxResults:10 Exact:false}
$ gor main.go 'http://localhost:12345/search?l=golang&l=progamming'
http-get: 200 OK
Search: {Labels:[golang progamming] MaxResults:10 Exact:false}
$ gor main.go 'http://localhost:12345/search?l=golang&l=progamming&max=100'
http-get: 200 OK
Search: {Labels:[golang progamming] MaxResults:100 Exact:false}
$ gor main.go 'http://localhost:12345/search?l=golang&l=progamming&x=true'
http-get: 200 OK
Search: {Labels:[golang progamming] MaxResults:10 Exact:true}
$ gor main.go 'http://localhost:12345/search?q=hello&x=123'
http-get: 400 Bad Request
x: strconv.ParseBool: parsing "123": invalid syntax
$ gor main.go 'http://localhost:12345/search?q=hello&max=lots'
http-get: 400 Bad Request
max: strconv.ParseInt: parsing "lots": invalid syntax
//!-output
*/
