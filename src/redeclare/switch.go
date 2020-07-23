package main

import "fmt"

var content string

//var content = "Python"

func main() {
	content := "Python"
	switch content {
	default:
		fmt.Println("Unknow language")
	case "Python":
		fmt.Println("An interpreted Language")
	case "Go":
		fmt.Println("A compiled language")
	}
}
