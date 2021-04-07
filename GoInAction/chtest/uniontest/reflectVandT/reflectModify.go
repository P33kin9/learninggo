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
	x := 2
	v := reflect.ValueOf(&x)
	fmt.Println(v)
	v.Elem().SetInt(100)
	fmt.Println(x)

	u := User{"张三", 30}
	v1 := reflect.ValueOf(&u.Name)
	v2 := reflect.ValueOf(&u.Age)
	fmt.Println(v1, v2)
	v1.Elem().SetString("里斯")
	v2.Elem().SetInt(22)
	fmt.Println(u)

}
