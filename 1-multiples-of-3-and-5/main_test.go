package main

import (
	"testing"
)

func TestNaive(t *testing.T) {
	res := naive(1000)
	if res != 233168 {
		t.Errorf("naive(1000) must be 233168. res=%v", res)
	}
}

func TestConstant(t *testing.T) {
	res := constant(1000)
	if res != 233168 {
		t.Errorf("constant(1000) must be 233168. res=%v", res)
	}
}
