package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bells Labs": {40.6834, -74.8920},
	"Google":     {37.42202, -122.0888384},
}

func main() {
	fmt.Println(m)
}
