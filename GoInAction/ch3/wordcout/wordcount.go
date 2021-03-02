package main

import (
	"fmt"
	"goination/ch3/words"
	"io/ioutil"
	"os"
)

// main is the entry point for the application.
func main() {
	filename := os.Args[1]

	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("There was an error opening the file:", err)
		return
	}

	text := string(contents)

	count := words.CountWords(text)
	fmt.Printf("There are %d words in your text. \n", count)
}
