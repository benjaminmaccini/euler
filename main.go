package main

import (
	"fmt"
	"math"

	"euler/helpers"
)

// Find the sum of all the multiples of 3 or 5 below 1000.
// Answer = 233168
func problem1Naive() int {
	sum := 0
	for i := 1; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}

// Positive 2x2 positive integer matrices that can be represented as the square of a 2x2 positive integer matrix
// If so their determinant is a perfect square.
// 1. Find all the matrices with perfect square determinants
func problem420() string {
	return "Not yet implemented"
}

// Find the smallest odd composite that cannot be written as the sum of a prime and twice a square
// k = p + 2*(a)^2
// (k - p)/2 = a^2
// So find the first k s.t. for all p's < k, (k - p)/2 is not a perfect square
func problem46(n float64) int {
	primes := helpers.FindPrimes(n)

	// Make into a map for efficient lookup
	primesMap := make(map[int]bool)
	for _, p := range primes {
		primesMap[p] = true
	}

	for k := 9; k < int(n); k += 2 {
		i := 0
		found := true
		for i < len(primes) && primes[i] <= k {
			if _, isPrime := primesMap[k]; isPrime {
				i++
				found = false
				continue
			}
			a := math.Sqrt(float64((k - primes[i]) / 2.0))

			// Check to see if it's an integer
			if a == math.Trunc(a) {
				found = false
			}
			i++
		}
		if found {
			return k
		}
	}

	return 0
}

func main() {
	fmt.Println(problem46(10000))
}
