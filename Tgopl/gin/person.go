package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	name string
	age  int
}

func (p Person) sayHello() {
	fmt.Println("Hello, my name is " + p.name + "and I am " + strconv.Itoa(p.age) + "years old ")
}
func main() {
	p := Person{
		"john",
		30,
	}

	fmt.Println(p.name)
	p.sayHello()
}
