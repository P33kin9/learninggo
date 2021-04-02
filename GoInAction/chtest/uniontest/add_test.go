package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	sum := Add(1, 2)
	if sum == 3 {
		t.Log("the result is Ok")
	} else {
		t.Fatal("the result is wrong")
	}
}
