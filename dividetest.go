package main

import (
	"errors"
	"fmt"
)

func divide(dividend int, divisor int) (result int, err error) {
	if divisor == 0 {
		err = errors.New("divisor by zero")
		return
	}
	result = dividend / divisor
	return
}

func main() {
	fmt.Println(divide(10, 5))
}

//divide(10 ,5)
