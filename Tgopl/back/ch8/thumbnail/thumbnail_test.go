// This file is just a place to put example code from the book.
// It does not actually run any code in cf/ch8/thumbnail.

package thumbnail_test

import (
	"log"
	"os"
	"sync"
	"testing"
	"tgopl/back/ch8/thumbnail"
)

// !+1
// makeThumbnails make thumbnails of the specified files.
func makeThumbnails(filenames []string) {
	for _, f := range filenames {
		if _, err := thumbnail.ImageFile(f); err != nil {
			log.Println(err)
		}
	}
}

//!-1

// !+2
// NOTE: incorrect!
func makeThumbnails2(filenames []string) {
	for _, f := range filenames {
		go thumbnail.ImageFile(f) // NOTE: ignoring errors
	}
}

//!-2

// !+3
// makeThumbnails3 makes thumbnails of the specified files in paralle.
func makeThumbnails3(filenames []string) {
	ch := make(chan struct{})
	for _, f := range filenames {
		go func(f string) {
			thumbnail.ImageFile(f) // NOTE: ignoring errors
			ch <- struct{}{}
		}(f)
	}

	// Wait for goroutines to complete.
	for range filenames {
		<-ch
	}
}

//!-3

// !+4
// makeThumbnails4 makes thumbnails for the specified files in parallel.
// It returns an error if any step failed.
func makeThumbnails4(filenames []string) error {
	errors := make(chan error)

	for _, f := range filenames {
		go func(f string) {
			_, err := thumbnail.ImageFile(f)
			errors <- err
		}(f)
	}

	for range filenames {
		if err := <-errors; err != nil {
			return err // NOTE: incorrect: goroutine leak!
		}
	}

	return nil
}

//!-4

// !+5
// makeThumbnails5 makes thumbnails for the specified files in paralls.
// It returns the generated file names in an arbitray order,
// or an error if any step failed.
func makeThumbnails5(filenames []string) (thumbnails []string, err error) {
	type item struct {
		thumbnail string
		err       error
	}

	ch := make(chan item, len(filenames))
	for _, f := range filenames {
		go func(f string) {
			var it item
			it.thumbnail, it.err = thumbnail.ImageFile(f)
			ch <- it
		}(f)
	}

	for range filenames {
		it := <-ch
		if it.err != nil {
			return nil, it.err
		}
		thumbnails = append(thumbnails, it.thumbnail)
	}

	return thumbnails, nil
}

//!-5

// !+6
// makeThumbnails6 makes thumbnails for each file received from the channel.
// It returns the number of bytes occupied by the files it creates.
func makeThumbnails6(filenames <-chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup // number of working goroutines
	for f := range filenames {
		wg.Add(1)
		// worker
		go func(f string) {
			defer wg.Done()
			thumb, err := thumbnail.ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thumb) // OK to ignore error
			sizes <- info.Size()
		}(f)
	}

	// closer
	go func() {
		wg.Wait()
		close(sizes)
	}()

	var total int64
	for size := range sizes {
		total += size
	}
	return total
}

//!-6

//  Test Functions

// TestMakeThumbnails tests makeThumbnails function
func TestMakeThumbnails(t *testing.T) {
	files := []string{"file1.jpg", "file2.jpg", "file3.jpg"}
	makeThumbnails(files)
}

// TestMakeThumbnails2 tests makeThumbnails2 function
func TestMakeThumbnails2(t *testing.T) {
	files := []string{"file1.jpg", "file2.jpg", "file3.jpg"}
	makeThumbnails2(files)
}

// TestMakeThumbnails3 tests makeThumbnails3 function
func TestMakeThumbnails3(t *testing.T) {
	files := []string{"file1.jpg", "file2.jpg", "file3.jpg"}
	makeThumbnails3(files)
}

// TestMakeThumbnails4 tests makeThumbnails4 function
func TestMakeThumbnails4(t *testing.T) {
	files := []string{"file1.jpg", "file2.jpg", "file3.jpg"}
	err := makeThumbnails4(files)
	if err != nil {
		t.Fatalf("makeThumbnails4 failed: %v", err)
	}
}

// TestMakeThumbnails5 tests makeThumbnails5 function
func TestMakeThumbnails5(t *testing.T) {
	files := []string{"file1.jpg", "file2.jpg", "file3.jpg"}
	thumbnails, err := makeThumbnails5(files)
	if err != nil {
		t.Fatalf("makeThumbnails5 failed: %v", err)
	}
	if len(thumbnails) != len(files) {
		t.Fatalf("makeThumbnails5 returned wrong number of thumbnails: got %d, want %d", len(thumbnails), len(files))
	}
}

// TestMakeThumbnails6 tests makeThumbnails6 function
func TestMakeThumbnails6(t *testing.T) {
	files := make(chan string, 3)
	files <- "file1.jpg"
	files <- "file2.jpg"
	files <- "file3.jpg"
	close(files)

	totalSize := makeThumbnails6(files)
	if totalSize == 0 {
		t.Fatalf("makeThumbnails6 returned zero size")
	}
}
