// The jpeg command reads a PNG image from the standard input
// and writes it as a JPEG image to the standard output.
package main

import (
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png"
	"io"
	"os"
)

func main() {
	if err := toJPEG(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
		os.Exit(1)
	}
}

func toJPEG(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}

//!-main

/*
//!+with
$ gor ../../ch3/mandelbrot/main.go | gor main.go > ~/Desktop/mandel.jpeg
Input format = png
//!-with

//!+without
$ gor ../../ch3/mandelbrot/main.go | gor main.go > ~/Desktop/mandel.jpeg
jpeg: image: unknown format
exit status 1
signal: broken pipe
//!-without
*/
