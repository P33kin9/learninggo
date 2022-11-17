package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var u1 user1
	var u2 user2
	var u3 user3
	var u4 user4
	var u5 user5
	var u6 user6

	fmt.Printf("u1 sizeof %d\n", unsafe.Sizeof(u1))
	fmt.Printf("u2 sizeof %d\n", unsafe.Sizeof(u2))
	fmt.Printf("u3 sizeof %d\n", unsafe.Sizeof(u3))
	fmt.Printf("u4 sizeof %d\n", unsafe.Sizeof(u4))
	fmt.Printf("u5 sizeof %d\n", unsafe.Sizeof(u5))
	fmt.Printf("u6 sizeof %d\n", unsafe.Sizeof(u6))
	//	fmt.Println(unsafe.Offsetof(u1.i))
	//	fmt.Println(unsafe.Offsetof(u1.j))
}

type user1 struct {
	b byte
	i int32
	j int64
}
type user2 struct {
	b byte
	j int64
	i int32
}
type user3 struct {
	i int32
	b byte
	j int64
}
type user4 struct {
	i int32
	j int64
	b byte
}
type user5 struct {
	j int64
	b byte
	i int32
}
type user6 struct {
	j int64
	i int32
	b byte
}
