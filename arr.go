package main

//
import "fmt"

//
//func main() {
//	var b = [...]int{1, 2, 3}
//	var c = [...]int{2: 3, 1: 2}
//	//var d = [...]int{1, 2, 4: 5, 6}
//	var d = [...]int{3: 5}
//	fmt.Printf("%d\n%d\n%d\n", b, c, d)
//}

var v interface{}

func main() {
	v = 4
	switch i := v.(type) {
	case string:
		fmt.Printf("The string is '%s'.\n", i)
	case int, uint, int8, uint8, int16, int32, uint32, int64, uint64:
		fmt.Printf("The integer is %T.\n", i)
	default:
		fmt.Printf("the type= %T.\n", i)
	}
}
