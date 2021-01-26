package popcount_test

import (
	"cf/ch2/popcount"
	"reflect"
	"testing"
)

// --Aternative implementations --
func assert(t *testing.T, expected, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("(expected, actual) = (%v, %v)\n", expected, actual)
	}
}

func TestPopCount(t *testing.T) {
	assert(t, 32, popcount.PopCount(0x123456789ABCDEF))
}

func TestPopCountByBitClear(t *testing.T) {
	assert(t, 32, popcount.PopCountByBitClear(0x123456789ABCDEF))
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(0x123456789ABCDEF)
	}
}

func BenchmarkPopCountByBitClear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountByBitClear(0x123456789ABCDEF)
	}
}
