package main

import "fmt"

func main() {
	x := 1
	p := &x // & memory address
	fmt.Println(p)
	*p = 2 //  * read point var value
	fmt.Println(x)
}
