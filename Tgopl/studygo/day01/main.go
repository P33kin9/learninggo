package main

import "fmt"

func Add(x int, y int) int {
	return x + y
}

func main(){
	fmt.Println(Add(42,3))
}