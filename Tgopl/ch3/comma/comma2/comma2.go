package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(comma2(-1234567.809))
}

func comma2(str float64) string {
	s := fmt.Sprintf("%.3f", str)
	var end string
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		end = s[dot:]
		s = s[:dot]
	}
	num := len(s)
	var buf bytes.Buffer
	j := 1
	for i := num - 1; i >= 0; i-- {
		buf.WriteByte(s[i])
		if j%3 == 0 && i != 0 {
			buf.WriteString(",")
		}
		j++
	}
	res := buf.String()
	var r bytes.Buffer
	for i := len(res) - 1; i >= 0; i-- {
		r.WriteByte(res[i])
	}
	r.WriteString(end)
	return r.String()
}
