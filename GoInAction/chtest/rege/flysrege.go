package main

import (
	"fmt"
	"regexp"
)

func main() {
	flysRegexp := regexp.MustCompile(`^http://www.flysnow.org/([\d]{4})/([\d]{2})/([\d]{2})/([\w-]+).html$`)
	params := flysRegexp.FindStringSubmatch("http://www.flysnow.org/2018/01/20/golang-goquery-examples-selecter.html")

	/*	for _, param := range params {
			fmt.Println(param)
		}
	*/
	fmt.Println("year: ", params[1])
	fmt.Println("month: ", params[2])
	fmt.Println("day: ", params[3])
	fmt.Println("title: ", params[4])
}
