package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.6543, -74.34498,
	},
	"Google": Vertex{
		37.98789, -122.873789,
	},
}

func main() {
	fmt.Println(m)
}
