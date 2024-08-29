package main

import (
	"fmt"

	"worktest/entities"
)

func main() {
	a := entities.Admin{
		Rights: 10,
	}

	a.Name = "Bill"
	a.Email = "bill@email.com"

	fmt.Printf("User: %v\n", a)
}
