package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		ch1 <- 1
	}()

	go func() {
		ch2 <- 2
	}()
	select {
	case x := <-ch1:
		fmt.Println(x)
	case y := <-ch2:
		fmt.Println(y)
	}
}
