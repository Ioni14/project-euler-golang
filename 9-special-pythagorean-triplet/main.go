package main

import (
	"fmt"
)

// a^2 + b^2 = c^2 and a + b + c = target
// <=> a^2 + b^2 = (target - a - b)^2
// <=> b = target * (target - 2*a) / (2 * (target - a)) and c = target - a - b
func naive(targetTripletSum uint) (targetProduct uint, triplet [3]uint) {
	for a := uint(1); a < targetTripletSum/2; a++ {
		numerator := targetTripletSum * (targetTripletSum - 2*a)
		denominator := 2 * (targetTripletSum - a)
		if numerator%denominator != 0 {
			continue
		}

		b := uint(numerator / denominator)
		if b < a {
			continue
		}
		c := targetTripletSum - a - b

		return a * b * c, [3]uint{a, b, c}
	}

	return 0, [3]uint{}
}

/*
source : https://projecteuler.net/problem=9

A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
a^2 + b^2 = c^2
For example, 32 + 42 = 9 + 16 = 25 = 52.
There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.
*/
func main() {
	res, _ := naive(1000)
	fmt.Println(res)
}
