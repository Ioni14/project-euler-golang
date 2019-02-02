package main

import (
	"testing"
)

func TestNaive(t *testing.T) {
	res := naive(10)
	if res != 17 {
		t.Errorf("naive(10) must be 17. res=%v", res)
	}

	res = naive(2000000)
	if res != 142913828922 {
		t.Errorf("naive(2000000) must be 142913828922. res=%v", res)
	}
}
