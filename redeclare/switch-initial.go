package main

import "fmt"
import "Strings"

func main() {
	content := " Go "
	switch lang := strings.TrimSpace(content); lang {
	default:
		fmt.Println("Unknow language")
	case "Python":
		fmt.Println("An interpreted language")
	case "Go":
		fmt.Println("A compiled language")
	}

}
