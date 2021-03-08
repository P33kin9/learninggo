// This sample code implement a simple ewb service.
package main

import (
	"goination/ch9/listing17/handlers"
	"log"
	"net/http"
)

// main is the entry point for the application.
func main() {
	handlers.Routes()

	log.Println("listener: Started : Listening on :4000")
	http.ListenAndServe(":4000", nil)
}
