package main

import "fmt"

//var number = 107
var number = 11

func main() {
	//number := 102
	if 100 < number {
		number++
	} else {
		number--
	}
	fmt.Printf("%d\n", number)
}

//if diff := 100 - number; 100 < diff {
//	number++
//} else {
//	number--
//}

//if diff := 100 - number; 100 < diff {
//	number++
//} else if 200 < diff {
//	number--
//} else {
//	number -=2
//}
