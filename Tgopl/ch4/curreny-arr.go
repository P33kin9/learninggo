package main

import "fmt"

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func main() {
	symbol := [...]string{USD: "$", EUR: "€", RMB: "¥"}
	fmt.Println(RMB, symbol[RMB]) // "3 ¥"
}
