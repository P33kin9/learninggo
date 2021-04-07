package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("Print reflect.TypeOf")
	u := User{"张三", 20}
	t := reflect.TypeOf(u)
	fmt.Println(t)

	fmt.Println("Print reflect.ValueOf")
	v := reflect.ValueOf(u)
	fmt.Println(v)

	fmt.Println("Print fmt.Printf " + "%" + "T and " + "%" + "V")
	fmt.Printf("%T\n", u)
	fmt.Printf("%v\n", u)

	fmt.Println("reflect.Value and reflect.Type")
	u1 := v.Interface().(User)
	fmt.Println(u1)

	t1 := v.Type()
	fmt.Println(t1)

	fmt.Println("Print underlying type of t.Kind()")
	fmt.Println(t.Kind())
}
