// The thumbnail command produces thumbnails of JPEG files
// whose names are provided on each line of the standard input.
//
// The "+build ignore" tag  excludes this file from the
// thumbnail package, but it can be compiled as a command and run like
// this:
//
// Run with:
//
//	$ go run $GOPATH/src/cf/ch8/thumbnail/main.go
//	foo.jpeg
//	^D
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"tgopl/back/ch8/thumbnail"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		thumb, err := thumbnail.ImageFile(input.Text())
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Println(thumb)
	}
	if err := input.Err(); err != nil {
		log.Fatal(err)
	}
}
