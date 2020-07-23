package main

import "fmt"

func outerFunc() {
	defer fmt.Println("The function over before execution")
	fmt.Println("The first execution")
}

func main() {
	outerFunc()
}
