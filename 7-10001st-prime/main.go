package main

import (
	"fmt"
)

var computedPrimes = []uint64{2, 3}

func IsComputedPrimeNumber(n uint64) bool {
	if n <= 1 {
		return false
	}
	for _, prime := range computedPrimes {
		if n == prime {
			return true
		}
		if n % prime == 0 {
			return false
		}
	}
	computedPrimes = append(computedPrimes, n)

	return true
}

func naive(max int) (lastPrime uint64) {

	countComputedPrimes := len(computedPrimes)

	i := computedPrimes[countComputedPrimes - 1] + 2
	for countComputedPrimes < max {
		if IsComputedPrimeNumber(i) {
			countComputedPrimes++
		}

		i += 2
	}

	return i - 2
}

/*
source : https://projecteuler.net/problem=7

By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
What is the 10 001st prime number?
*/
func main() {
	res := naive(10001)
	fmt.Println(res)
}
