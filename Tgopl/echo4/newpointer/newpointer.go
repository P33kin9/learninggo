package main

import "fmt"

func main() {
	p := new(int)
	q := new(int)
	fmt.Printf("p = %d, q = %d\n", &p, &q)
	fmt.Println(p == q)
}
