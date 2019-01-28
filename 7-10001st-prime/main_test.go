package main

import (
	"testing"
)

func TestNaive(t *testing.T) {
	res := naive(6)
	if res != 13 {
		t.Errorf("naive(6) must be 13. res=%v", res)
	}

	res = naive(10001)
	if res != 104743 {
		t.Errorf("naive(10001) must be 104743. res=%v", res)
	}
}
