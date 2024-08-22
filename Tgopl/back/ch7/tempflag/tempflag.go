package main

import (
	"tgopl/back/ch7/tempconv"
	"flag"
	"fmt"
)

//!+
//var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

var tempF = tempconv.FahrenheitFlag("temp", -18.0, "the temperature")

func main() {
	flag.Parse()
	//	fmt.Println(*temp)
	fmt.Println(*tempF)
}
