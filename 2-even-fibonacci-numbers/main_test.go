package main

import (
	"testing"
)

func TestNaive(t *testing.T) {
	res := naive(4000000)
	if res != 4613732 {
		t.Errorf("naive(4000000) must be 4613732. res=%v", res)
	}
}
