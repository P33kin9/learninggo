package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	//exercise 1.1
	// for i:=1 modify to i :=0
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
