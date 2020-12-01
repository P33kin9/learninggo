package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(compare("abec", "ecab"))
}

func compare(str1 string, str2 string) bool {
	num1 := strings.Count(str1, "")
	num2 := strings.Count(str2, "")
	if num2 > num1 {
		str1, str2 = str2, str1
	}
	var res bool
	for _, v := range str1 {
		res = false
		for _, sv := range str2 {
			if v == sv {
				res = true
			}
		}
		if !res {
			break
		}
	}
	return res
}
