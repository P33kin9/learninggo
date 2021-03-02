package main

import (
	_ "goination/ch2/sample/matchers"
	"goination/ch2/sample/search"
	"log"
	"os"
)

// init is called prior to main.
func init() {
	// Change the device for logging to stdout.
	log.SetOutput(os.Stdout)
}

// main is the entry point for the progam.
func main() {
	search.Run("[p|P]resident")
}
