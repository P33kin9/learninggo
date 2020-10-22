package main

import "fmt"

func main() {
	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("Recovered panic: %s\n", p)
		}
	}()
	myIndex := 4
	ia := [3]int{1, 2, 3}
	_ = ia[myIndex]
}

