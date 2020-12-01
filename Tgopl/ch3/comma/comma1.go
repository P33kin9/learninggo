package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("1234567890"))
}

func comma(s string) string {
	var newByte byte = ','
	n := len(s)
	buf := bytes.NewBuffer([]byte{})

	if n <= 3 {
		return s
	}

	for i := 0; i < n; i++ {

		if (n-i)%3 == 0 && i != 0 {
			buf.WriteByte(newByte)
		}
		buf.WriteByte(s[i])
	}
	return buf.String()
}
