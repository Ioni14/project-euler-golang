package main

import (
	"fmt"
)

// uses remarkable identity (a+b)^2 = a^2 + 2ab + b^2
// (a + b + c)^2 - (a^2 + b^2 + c^2)
// = a^2 + 2ab + 2ac + b^2 + 2bc + c^2 - (a^2 + b^2 + c^2)
// = 2ab + 2ac + 2bc
// = 2*(ab + ac + bc)
func naive(max int) (difference int64) {
	var res int64 = 0
	for i := 1; i <= max - 1; i++ {
		for j := i + 1; j <= max; j++ {
			res += int64(i * j)
		}
	}

	return 2 * res
}

/*
source : https://projecteuler.net/problem=6

The sum of the squares of the first ten natural numbers is,
1^2 + 2^2 + ... + 10^2 = 385
The square of the sum of the first ten natural numbers is,
(1 + 2 + ... + 10)^2 = 552 = 3025
Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.
Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
*/
func main() {
	res := naive(100)
	fmt.Println(res)
}
