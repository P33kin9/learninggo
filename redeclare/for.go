package main

import "fmt"

func main() {
	var m = 1
	for m < 50 {
		m *= 3
		fmt.Printf("The m is %d\n", m)
	}
	fmt.Printf("The m is %d\n", m)
}
