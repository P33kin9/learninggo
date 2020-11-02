package main

import "fmt"

type I interface {
	m()
}

type MyI1 I
type MyI2 = I

type MyInt int

func (i MyInt) m() {
	fmt.Println("MyInt.m")
}

func main() {
	var i I = MyInt(0)
	var i1 MyI1 = MyInt(0)
	var i2 MyI2 = MyInt(0)

	i = i1
	i = i2
	i1 = i2
	i1 = i
	i2 = i1
}
