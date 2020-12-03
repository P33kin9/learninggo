package main

import (
	"fmt"
	. "net"
)

func IsUp(v Flags) bool     { return v&FlagUp == FlagUp }
func TurnDown(v *Flags)     { *v &^= FlagUp }
func SetBroadcast(v *Flags) { *v |= FlagBroadcast }
func IsCast(v Flags) bool   { return v&(FlagBroadcast|FlagMulticast) != 0 }

func main() {
	var v Flags = FlagMulticast | FlagUp
	fmt.Printf("%b %t\n", v, IsUp(v)) // "10001 true"
	TurnDown(&v)
	fmt.Printf("%b %t\n", v, IsUp(v)) // "1000 false"
	SetBroadcast(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))   //"10010 true"
	fmt.Printf("%b %t\n", v, IsCast(v)) //"10010 true"
}
