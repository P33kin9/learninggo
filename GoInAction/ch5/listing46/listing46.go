// Sample program to show how you can't always get the
// address of a value.
package main

import "fmt"

// duration is a type with a ase type of int.
type duration int

// format pretty-prints the duration value.
func (d duration) pretty() string {
	return fmt.Sprintf("Duration: %d", &d) // modify *duration and *d => &d
}

// main is the entry point for the application.
func main() {
	duration(42).pretty()
	fmt.Printf("value address: %s", duration(42).pretty()) // add the line and display value
	// ./listing46.go:17:14: cannot call pointer method pretty on durtion(42)
	// ./listing46.go:17:13: cannot take the address of durtion(42)

}
