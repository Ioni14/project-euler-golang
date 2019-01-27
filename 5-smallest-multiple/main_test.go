package main

import (
	"testing"
)

func TestNaive(t *testing.T) {
	res := naive(10)
	if res != 2520 {
		t.Errorf("naive(10) must be 2520. res=%v", res)
	}

	res = naive(20)
	if res != 232792560 {
		t.Errorf("naive(20) must be 232792560. res=%v", res)
	}
}
