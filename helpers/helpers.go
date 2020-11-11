package helpers

import (
	"math"
)

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

// https://stackoverflow.com/questions/30226438/generate-all-permutations-in-go
func Permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func Factorial(n int) (result int) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}
