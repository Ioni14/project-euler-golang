package main

import (
	"testing"
)

func TestNaive(t *testing.T) {
	res := naive(10)
	if res != 2640 {
		t.Errorf("naive(10) must be 2640. res=%v", res)
	}

	res = naive(100)
	if res != 25164150 {
		t.Errorf("naive(100) must be 25164150. res=%v", res)
	}
}
