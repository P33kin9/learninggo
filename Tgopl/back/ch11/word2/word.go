// Package word provides utilities for word games.
//!+
package work

import (
	"unicode"
)

// IsPalindrome reports whether s reads the same forward and backword.
// Letter case is ignored, as are non-letters.
func IsPalindrome(s string) bool {
	//	var letters []rune
	// BenchmarkIsPalinfrome performance improve 35%.

	letters := make([]rune, 0, len(s))
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}
	for i := range letters {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}

//!-
