package main

import (
	"testing"
)

func TestNaive(t *testing.T) {
	res, _ := naive(12)
	if res != 60 {
		t.Errorf("naive(12) must be 60. res=%v", res)
	}

	res, _ = naive(1000)
	if res != 31875000 {
		t.Errorf("naive(1000) must be 31875000. res=%v", res)
	}
}
