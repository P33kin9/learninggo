// Sample program to show how the program can access a value
// of an unexported identifier from anther package.
package main

import (
	"fmt"
	"goination/ch5/listing68/counters"
)

// main is entry point for the application.
func main() {
	// Create a variable of the unexported type using the exported
	// New function from the package counters.
	counter := counters.New(10)

	fmt.Printf("Counter: %d\n", counter)

}
