package format_test

import (
	"cf/ch12/format"
	"fmt"
	"testing"
	"time"
)

func Test(t *testing.T) {
	// The pointer values are just examples, and may vary from run to run.
	//!+time
	var x int64 = 1
	var d time.Duration = 1 * time.Nanosecond
	fmt.Println(format.Any(x))                  // "1"
	fmt.Println(format.Any(d))                  // "1"
	fmt.Println(format.Any([]int64{x}))         // "[]int64 0xc0000a6080"
	fmt.Println(format.Any([]time.Duration{d})) // "[]time.Duration 0xc0000a6088"
	//!-time
}
