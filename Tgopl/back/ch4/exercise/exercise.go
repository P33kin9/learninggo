package main

import (
	"fmt"
	"unicode"
)

func main() {
	testArr := [5]int{0, 1, 2, 3, 4}
	// reverse(&testArr)
	fmt.Println(testArr)
	a := rotate(testArr[:], 2)
	fmt.Println(a)

	b := []string{"tao", "taoshihan", "shi", "shi", "han"}
	emptyString(b)
	d := []byte("abc bcd wer sdsd taoshihan     de")
	e := emptyString2(d)

	fmt.Println(string(e))
	f := []byte("abc bcd wer sdsd  taoshiban       de ")
	reverse1(f)
	fmt.Println(string(f))
}

func reverse(s *[5]int) {
	i, j := 0, len(*s)-1
	for i < j {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
		i += 1
		j -= 1
	}
}

func rotate(s []int, r int) []int {
	lens := len(s)

	res := make([]int, lens)
	for i := 0; i < lens; i++ {
		index := i + r
		if index >= lens {
			index = index - lens
		}
		res[i] = s[index]
	}
	return res
}

func emptyString(s []string) []string {
	i := 0
	index := 0
	num := len(s)
	for _, v := range s {
		index = i + 1
		if index >= num {
			break
		}
		if v != s[index] {
			s[i] = v
			i++
		}
	}
	fmt.Println(s[:i])
	return s[:i]
}

func emptyString2(s []byte) []byte {
	index := 0
	num := len(s)
	for i := 0; i < num; i++ {
		index = i + 1
		num = len(s)
		if index >= num {
			break
		}
		if unicode.IsSpace(rune(s[i])) && unicode.IsSpace(rune(s[index])) {
			copy(s[i:], s[index:])
			s = s[:len(s)-1]
			i--
		}
	}
	return s
}

func reverse1(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
