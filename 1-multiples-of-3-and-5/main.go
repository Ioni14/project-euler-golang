package main

import (
	"fmt"
	"math"
)

func naive(n int) (cc int) {
	for i := 3; i < n; i++ {
		if i%3 == 0 || i%5 == 0 {
			cc += i
		}
	}
	return
}

// use the sum of integers from 1 to n
// don't forget to remove the duplicates (one for 3 and one for 5, so remove one for 3*5=15)
func constant(n int) (cc int) {
	k3 := math.Floor(float64(n - 1) / 3)
	k5 := math.Floor(float64(n - 1) / 5)
	k15 := math.Floor(float64(n - 1) / 15)
	cc = int(3 * k3 * (k3 + 1) / 2)
	cc += int(5 * k5 * (k5 + 1) / 2)
	cc -= int(15 * k15 * (k15 + 1) / 2)
	return
}

/*
source : https://projecteuler.net/problem=1

If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
Find the sum of all the multiples of 3 or 5 below 1000.
*/
func main() {
	//res := naive(1000)
	res := constant(1000)
	fmt.Println(res)
}
