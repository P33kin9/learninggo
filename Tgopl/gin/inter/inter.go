package main

import "fmt"

type Shape1 interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
func calculationArea(s Shape1) {
	fmt.Println("Area: ", s.Area())
}

func main() {

	rect := Rectangle{
		Width:  5,
		Height: 3,
	}
	calculationArea(rect)
}
