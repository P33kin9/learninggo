package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	var c WordsCounter
	fmt.Fprintf(&c, "hello world 123")
	fmt.Println(c)
}

/*
练习 7.1： 使用来自ByteCounter的思路，实现一个针对单词和行数的计数器。你会发现bufio.ScanWords非常的有用。
*/

//declare format
type WordsCounter int

func (w *WordsCounter) Write(p []byte) (int, error) {
	s := strings.NewReader(string(p))
	bs := bufio.NewScanner(s)
	bs.Split(bufio.ScanWords)

	sum := 0
	for bs.Scan() {
		sum++
	}

	*w = WordsCounter(sum)
	return sum, nil
}
