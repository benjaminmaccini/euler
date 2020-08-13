package helpers

import "math"

// Sieve of Eratosthenes
func FindPrimes(n float64) []int {
	all := make([]bool, int(n))
	for i := 2; i < int(n); i++ {
		all[i] = true
	}

	nSqrt := int(math.Sqrt(n)) + 1

	for p := 2; p < nSqrt; p++ {
		if all[p] {
			j := p * p
			for j < int(n) {
				all[j] = false
				j += p
			}
		}
	}

	primes := []int{}
	for i, e := range all {
		if e {
			primes = append(primes, i)
		}
	}

	return primes
}
