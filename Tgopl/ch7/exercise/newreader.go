package main

import (
	"fmt"
	"io"
)

type StringReader struct {
	data    string
	current int
}

func (sr *StringReader) Read(b []byte) (n int, err error) {
	if len(b) == 0 {
		return 0, nil
	}

	n = copy(b, sr.data[sr.current:])
	if sr.current += n; sr.current >= len(sr.data) {
		err = io.EOF
	}
	return
}

func NewReader(in string) *StringReader {
	sr := new(StringReader)
	sr.data = in
	return sr
}

func main() {
	str := "Hello World"
	sr := NewReader(str)
	data := make([]byte, 10)
	n, err := sr.Read(data[:0])
	for err == nil {
		n, err = sr.Read(data)
		fmt.Println(n, string(data[0:n]))
	}
	//output:
	// 10 Hello Worl
	// 1 d
}
