package main

import "fmt"

type Point struct{ X, Y int }

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	var w Wheel
	//	w = Wheel{Circle{Point{X: 8, Y: 8}, 5}, 20}
	w = Wheel{
		Circle: Circle{
			Point:  Point{X: 7, Y: 7},
			Radius: 5,
		},
		Spokes: 20, // NOTE: trailing comma necessary here (and Radius)
	}

	fmt.Printf("%#v\n", w)

	w.X = 46

	fmt.Printf("%#v\n", w)
}
