package main

import (
	"bufio"
	"fmt"
	"strings"
)

type LinesCounter int

func (l *LinesCounter) Writer(p []byte) (int, error) {
	lines := strings.NewReader(string(p))
	bs := bufio.NewScanner(lines)
	bs.Split(bufio.ScanLines)

	sum := 0
	for bs.Scan() {
		sum++
	}

	*l += LinesCounter(sum)
	return sum, nil
}

func main() {
	var l LinesCounter
	lines := "a\nb\nc\n"
	lines1 := "aa\nbb\ncc\n"
	l.Writer([]byte(lines))
	l.Writer([]byte(lines1))
	fmt.Println(l)
}
