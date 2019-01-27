package main

import (
	"fmt"
	"math"
)

func isPalindrome(n int) bool {
	maxPower := int(math.Floor(math.Log10(float64(n))))

	for cc := 0; cc < (maxPower + 1) / 2; cc++ {
		if int(n / int(math.Pow10(maxPower - cc))) % 10 != int(n / int(math.Pow10(cc))) % 10 {
			return false
		}
	}

	return true
}

func naive(min int) (maxPalindrome int) {
	maxPalindrome = 0
	for a := min; a <= min * 10 - 1; a++ {
		for b := a; b <= min * 10 - 1; b++ {
			n := a * b

			if isPalindrome(n) && n > maxPalindrome {
				maxPalindrome = n
			}
		}
	}

	return maxPalindrome
}

/*
source : https://projecteuler.net/problem=4

A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
Find the largest palindrome made from the product of two 3-digit numbers.
*/
func main() {
	res := naive(100)
	fmt.Println(res)
}
