package main

import "fmt"

var v = "1, 2, 3"

func main() {
	fmt.Printf("%v\n", v)
	v := []int{1, 2, 3}
	fmt.Printf("%v\n", v)
	if v != nil {
		var v = 123
		fmt.Printf("%v\n", v)
	}
}
