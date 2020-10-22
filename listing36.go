package main

import "fmt"

type duration int

func (d duration) pretty() string {
	return fmt.Sprintf("Duratrion %d", d)
}

func main() {
	duration(42).pretty()
	//	a := duration(42)
	//	fmt.Printf("fmt %X", a.pretty())
}
