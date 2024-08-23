package main

import (
	"fmt"
	"os"
	"time"
)

// exercise 1.3
// add time.Now() and time.Since(start).Seconds()
func main() {
	start := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	fmt.Println(time.Since(start).Microseconds())
}

// exercise 1.2
// // func main(){
// 	for i, arg := range os.Args[1:]{
// 		fmt.Println(i,arg)
// 	}
// }
