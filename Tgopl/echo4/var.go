package main

import "fmt"

var global *int

func f() {
	var x int
	x = 3
	global = &x
	fmt.Printf("f() local global= %x\n", global)
}

func g() {
	y := new(int)
	*y = 2
	fmt.Printf("*y = %d, &y = %d\n", *y, &y)
}

func main() {
	f()
	g()
	fmt.Printf("main global = %x\n", *global)
}
