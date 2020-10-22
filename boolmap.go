package main

import "fmt"

func wathc(person string) {
	attended := map[string]bool{
		"Ann": true,
		"Joe": true,
	}
	if attended[person] {
		fmt.Println(person, "meeting")
	} else {
		fmt.Println("Nobody")
	}
}

func main() {
	wathc("fonh")
}
