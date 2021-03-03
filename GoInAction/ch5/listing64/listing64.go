// Sample program to show how the program can't access an
// unexported identifier from another package.
package main

import (
	"fmt"
	"goination/ch5/listing64/counters"
)

//main is the entry point for the application.
func main() {
	// Create a variable fo the unexported type and initialize
	// the value to 10.
	counter := counters.alertCounter(10)

	// ./listing64.go:14:13: cannot refer to unexported name counters.alertCounter
	fmt.Printf("Counter: %d\n", counter)
}
