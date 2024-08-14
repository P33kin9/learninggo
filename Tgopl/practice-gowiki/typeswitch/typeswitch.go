package main

import "fmt"

func funcitonOfSomeType() bool {
	return true
}

func main() {
	var t interface{}
	t = funcitonOfSomeType()

	switch t := t.(type) {
	case bool:
		fmt.Printf("boolean %t\n", t)
	case int:
		fmt.Printf("integer %d\n", t)
	case *bool:
		fmt.Printf("pointer to boolean %t\n", *t)
	case *int:
		fmt.Printf("pointer to integer %d\n", *t)
	default:
		fmt.Printf("unexpected type %T\n", t)
	}
}
