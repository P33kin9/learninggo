package main

import (
	"Strings"
	"fmt"
)

func main() {
	content := "Go "
	switch lang := strings.TrimSpace(content); lang {
	case "Ruby":
		fallthrough
	case "Python":
		fmt.Println("An interpreted language")
	case "Go":
		fmt.Println("A compiled language")
	default:
		fmt.Println("Unknow language")
	}
}
