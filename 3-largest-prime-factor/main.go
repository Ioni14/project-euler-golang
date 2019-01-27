package main

import (
	"fmt"
	"math"
)

var computedPrimes = []int{2, 3}//, 5, 7, 11, 13}

func GetAPrimeFactor(n int64) int64 {
	if n % 2 == 0 {
		return 2
	}
	for i := int64(3); i <= int64(math.Sqrt(float64(n))); i+=2 {

		// is i a prime number ?
		addToPrimes := true
		for _, prime := range computedPrimes {
			if i == int64(prime) {
				if n % i == 0 {
					return i // we found a prime factor (i) of n
				}
				addToPrimes = false // already in
				break
			}
			if i % int64(prime) == 0 {
				// divisible by prime : it is not a prime number
				addToPrimes = false // not a prime
				break
			}
		}
		if addToPrimes  {
			computedPrimes = append(computedPrimes, int(i))
			if n % i == 0 {
				return i
			}
		}
	}

	return n
}

func naive(n int64) (maxPrimeNumber int64) {
	maxPrimeNumber = 1
	for n > 1 {
		p := GetAPrimeFactor(n)
		n /= int64(p)
		if p > maxPrimeNumber {
			maxPrimeNumber = p
		}
	}
	return
}

/*
source : https://projecteuler.net/problem=3

The prime factors of 13195 are 5, 7, 13 and 29.
What is the largest prime factor of the number 600851475143 ?
*/
func main() {
	res := naive(600851475143)
	fmt.Println(res)
}
