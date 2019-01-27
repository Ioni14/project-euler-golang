package main

import (
	"testing"
)

func TestNaive(t *testing.T) {
	res := naive(13195)
	if res != 29 {
		t.Errorf("naive(13195) must be 29. res=%v", res)
	}

	res = naive(600851475143)
	if res != 6857 {
		t.Errorf("naive(600851475143) must be 6857. res=%v", res)
	}
}
