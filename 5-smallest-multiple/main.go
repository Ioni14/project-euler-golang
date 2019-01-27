package main

import (
	"fmt"
	"math"
)

var primes = []int{2,3,5,7,11,13,17,19}

func naive(max int) (multiple int64) {
	multiple = 1
	for _, prime := range primes {
		if prime > max {
			break
		}
		maxPrimePower := 0
		for int(math.Pow(float64(prime), float64(maxPrimePower))) <= max {
			maxPrimePower++
		}
		multiple *= int64(math.Pow(float64(prime), float64(maxPrimePower - 1)))
	}
	return
}

/*
source : https://projecteuler.net/problem=5

2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/
func main() {
	res := naive(20)
	fmt.Println(res)
}
