// Calculates the number of different bits in two SHA256 Hash codes
// By default, it prints the SHA256 encoding of standard input and
// supports flag customization from the command line to output the
// SHA384 or Hasche Algorithm.
package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"strings"
)

// Command line flag
var hashMethod = flag.String("s", "sha256", "请输入哈希算法")

func main() {
	flag.Parse()
	num := compareSha256("x", "X")
	fmt.Println(num)

	// Flag's on the command line.
	printHash(strings.ToUpper(*hashMethod), "x")
}

func compareSha256(str1 string, str2 string) int {
	a := sha256.Sum256([]byte(str1))
	b := sha256.Sum256([]byte(str2))
	num := 0
	// Loop Byte array
	for i := 0; i < len(a); i++ {
		// One byte equals eight bits, the shift operation, gets each bit.
		for m := 1; m <= 8; m++ {
			// compare bits
			if (a[i] >> uint(m)) != (b[i] >> uint(m)) {
				num++
			}
		}
	}
	return num
}

func printHash(flag string, str string) {
	if flag == "SHA256" {
		fmt.Printf("%x\n", sha256.Sum256([]byte(str)))
		return
	}
	if flag == "SHA512" {
		fmt.Printf("%x\n", sha512.Sum512([]byte(str)))
		return
	}
	if flag == "SHA384" {
		fmt.Printf("%x\n", sha512.Sum384([]byte(str)))
		return
	}
}
