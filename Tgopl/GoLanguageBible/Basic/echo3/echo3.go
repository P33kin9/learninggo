package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// exercise 1.3
// add time.Now() and time.Since(start).Seconds()
func main() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(time.Since(start).Microseconds())
}
