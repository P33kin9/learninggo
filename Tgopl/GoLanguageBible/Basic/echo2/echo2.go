package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
// exercise 1.2
// // func main(){
// 	for i, arg := range os.Args[1:]{
// 		fmt.Println(i,arg)
// 	}
// }