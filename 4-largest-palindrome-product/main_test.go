package main

import (
	"testing"
)

func TestNaive(t *testing.T) {
	res := naive(10)
	if res != 9009 {
		t.Errorf("naive(10) must be 9009. res=%v", res)
	}

	res = naive(100)
	if res != 906609 {
		t.Errorf("naive(100) must be 906609. res=%v", res)
	}
}
