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
		if n%prime == 0 {
			return false
		}
	}
	computedPrimes = append(computedPrimes, n)

	return true
}

func naive(max uint64) (sumPrimes uint64) {
	countComputedPrimes := len(computedPrimes)

	i := computedPrimes[countComputedPrimes-1] + 2
	for _, computedPrime := range computedPrimes {
		sumPrimes += computedPrime
	}
	for i < max {
		if IsComputedPrimeNumber(i) {
			countComputedPrimes++
			sumPrimes += i
		}

		i += 2
	}
	return
}

/*
source : https://projecteuler.net/problem=10

The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
Find the sum of all the primes below two million.
*/
func main() {
	res := naive(2000000)
	fmt.Println(res)
}
