package main

import "fmt"

func main() {
	var number int
	for i := 0; i < 100; i++ {
		number++
	}
	fmt.Printf("The number is %d\n", number)

	var j uint = 1
	for ; j%5 != 0; j *= 3 {
		number++
	}
	fmt.Printf("The number is %d\n", number)

	for k := 0; k%5 != 0; {
		k *= 3
		number++
	}
	fmt.Printf("The number is %d\n", number)
}
