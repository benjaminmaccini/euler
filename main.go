package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"euler/helpers"
)

// Find the sum of all the multiples of 3 or 5 below 1000.
// Status: Complete
func problem1Naive() int {
	sum := 0
	for i := 1; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}

// Find the millionth permutation (i.e. 012 < 021 < 102 < ...) for the digits 0-9
// Status: Complete
func problem24() string {
	digits := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	count := 1000000
	i := 0

	final := ""

	for len(digits) > 0 {
		step := helpers.Factorial(len(digits) - 1)

		fmt.Printf("step: %d, count %d\n", step, count)
		if (count - step) <= 0 {
			final += digits[i]
			if i < len(digits)-1 {
				digits = append(digits[0:i], digits[i+1:]...)
			} else {
				digits = digits[0:i]
			}
			fmt.Printf("final: %s, remaining %v\n", final, digits)
			i = 0
		} else {
			count -= step
			i++
		}

	}

	return final
}

// How many circular primes are there below one million?
// Status: Complete
func problem35(n float64) int {
	primes := helpers.FindPrimes(n)

	// Start like this to account for the 2
	count := 1

	for i, p := range primes {
		ps := strconv.Itoa(p)
		tally := 1
		for j := 0; j < len(ps); j++ {
			// Cycle the number
			ps = ps[1:] + string(ps[0])

			// Just let it blow up if there is an error
			checkMe, _ := strconv.Atoi(ps)

			// Because we add the total count we shouldn't need to traverse the whole list again
			for _, pp := range primes[i+1:] {
				if checkMe == pp {
					tally += 1
					break
				}
			}
		}

		// If the tally matches add the tally to the total count
		if tally == len(ps) {
			count += tally
		}
	}

	return count
}

// Find the smallest odd composite that cannot be written as the sum of a prime and twice a square
// k = p + 2*(a)^2
// (k - p)/2 = a^2
// So find the first k s.t. for all p's < k, (k - p)/2 is not a perfect square
// Status: Complete
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

// Find the smallest x s.t. 2x, 3x, 4x, 5x, and 6x all have the same digits (in different orders)
// Status: Complete
func problem52() int {
	for i := 1; i <= 100000000; i++ {
		s := strconv.Itoa(i)
		s2 := strconv.Itoa(i * 2)
		s3 := strconv.Itoa(i * 3)
		s4 := strconv.Itoa(i * 4)
		s5 := strconv.Itoa(i * 5)
		s6 := strconv.Itoa(i * 6)

		for _, c := range s {
			s2 = strings.Replace(s2, string(c), "", 1)
			s3 = strings.Replace(s3, string(c), "", 1)
			s4 = strings.Replace(s4, string(c), "", 1)
			s5 = strings.Replace(s5, string(c), "", 1)
			s6 = strings.Replace(s6, string(c), "", 1)
		}

		if "" == s2 && "" == s3 && "" == s4 && "" == s5 && "" == s6 {
			return i
		}
	}

	return 0
}

// Find the length of the longest amicable chain with all elements < 1000000
// Don't need to search for anything > 500000
// Status: Incomplete
func problem95() int {
	primes := helpers.FindPrimes(50.0)

	// Find sums, map to the number
	sums := make(map[int]int)
	start := time.Now()
	for i := 1; i < 37; i++ {
		sum := 1 - i // This is to account for prime factoring method below
		// Find the prime factors, with their exponents
		primeFactorsMap := map[int]int{}
		j := i
		for {
			if j == 0 || j == 1 {
				break
			}
			for _, p := range primes {
				k := j % p
				if k == 0 {
					if _, exists := primeFactorsMap[p]; exists {
						primeFactorsMap[p] += 1
					} else {
						primeFactorsMap[p] = 1
					}
					j /= p
					break
				}
			}
		}

		// Convert to a list for permutation
		primeFactors := []int{}
		for p, v := range primeFactorsMap {
			for k := 0; k < v; k++ {
				primeFactors = append(primeFactors, p)
			}
		}

		divisorPrimeFactors := helpers.Permutations(sort.IntSlice(primeFactors))
		divisors := []int{}
		for _, dFactors := range divisorPrimeFactors {
			d := 1
			for _, f := range dFactors {
				d *= f
			}
		}

		for _, d := range divisors {
			sum += d
		}
		sums[i] = sum
	}
	t := time.Now().Sub(start)
	fmt.Printf("Find sums time: %s\n", t.String())

	maxLen := 1

	// Find chains
	for n, s := range sums {
		chainOrdered := []int{n, s}
		chain := map[int]int{n: s}
		ns := s
		for {
			_, exists := chain[ns]
			if exists {
				break
			} else {
				nns, _ := sums[ns]
				chain[ns] = nns
				chainOrdered = append(chainOrdered, nns)
				ns = nns
			}
		}

		// Check to see if the chain is amicable
		if chainOrdered[0] == chainOrdered[len(chainOrdered)-1] {
			// Count the length
			chainLen := len(chainOrdered) - 1
			if chainLen > maxLen {
				fmt.Printf("New longest chain found len:%d chainOrdered:%#v chain:%#v\n", chainLen, chainOrdered, chain)
				maxLen = chainLen
			}
		}
	}

	return maxLen
}

// Find the minimal connected network (minimal spanning tree) with the largest savings from a 40 node network
// Provided in p107_network.txt or p107_test.txt
// https://en.wikipedia.org/wiki/Minimum_spanning_tree
// Using Prim's Algorithm
// Status: Complete
func problem107() int {
	file, err := os.Open("p107_network.txt")
	if err != nil {
		fmt.Println(err)
	}

	r := csv.NewReader(file)
	networkMat, _ := r.ReadAll()

	network := make(map[int]map[int]int)
	networkValue := 0

	// Wrangle into a weighted adjacency list like {1: {1: 1, 2: 2}}
	for r, row := range networkMat {
		for c, val := range row {
			if c == r || val == "-" {
				continue
			}
			weight, _ := strconv.ParseInt(val, 10, 64)
			rel, exists := network[r]
			if !exists {
				rel = make(map[int]int)
				network[r] = rel
			}
			rel[c] = int(weight)
			networkValue += int(weight)
		}
	}

	networkValue /= 2 // Account for double visits

	// Begin Prim's
	costs := make(map[int]int)     // Map the vertices to their cheapest edge
	edges := make(map[int]int)     // Map between vertices with the cheapest edge
	vertices := make(map[int]bool) // Really just a set
	mst := make(map[int]int)       // The final tree

	for node, _ := range network {
		costs[node] = 10000
		edges[node] = -1
		vertices[node] = true
	}

	// debugLabels := "ABCDEFG"

	// Find and remove the minimum vertex
	for len(vertices) > 0 {
		minVertexCost := 10000
		minVertex := -1
		for v, _ := range vertices {
			if costs[v] < minVertexCost {
				minVertexCost = costs[v]
				minVertex = v
			}
		}
		if minVertex == -1 {
			minVertex = rand.Intn(len(network))
		}
		// fmt.Printf("minVert=%s, cost=%d\n", string(debugLabels[minVertex]), minVertexCost)
		delete(vertices, minVertex)

		// Add to the mst
		if edges[minVertex] != -1 {
			mst[minVertex] = edges[minVertex]
		}

		for edge, cost := range network[minVertex] {
			_, in := vertices[edge]
			if in && cost < costs[edge] {
				costs[edge] = cost
				edges[edge] = minVertex
			}
		}
	}

	treeValue := 0
	for node, _ := range mst {
		treeValue += costs[node]
	}

	reduction := networkValue - treeValue

	fmt.Println(mst)
	fmt.Printf("total=%d, tree=%d, loss=%d", networkValue, treeValue, reduction)

	return reduction
}

type AdditionTreeNode struct {
	value    int
	children []*AdditionTreeNode
	parent   *AdditionTreeNode
}

// Recursively build a tree
// limit is the max tree depth, initial value being 0
// Keep track of the lowest depth and use that as the minimal chain length, we don't actually have to
// know the chain
func buildAdditionTree(node *AdditionTreeNode, chain []int, limit int, depths map[int]int) {

	if depth, exists := depths[node.value]; !exists {
		depths[node.value] = limit
	} else if limit < depth {
		depths[node.value] = limit
	}

	// Kill the tree if the depth is >11, this was determined via Flammenkamp's a(n) <= 9/log_2(n)*log_2(v(n))
	// where v(n) is the population count of the binary representation
	limit += 1
	if limit <= 11 {
		for _, a := range chain {
			childValue := node.value + a
			nextChain := append(chain, childValue)
			grandchildren := make([]*AdditionTreeNode, 0)
			child := AdditionTreeNode{
				value:    childValue,
				children: grandchildren,
				parent:   node,
			}
			node.children = append(node.children, &child)
			buildAdditionTree(&child, nextChain, limit, depths)
		}
	}
}

// Find the sum of the minimal addition chains for all n <= 200
// Thurber's Algorithm (not used since I just need to know the depths): https://pdfs.semanticscholar.org/6e33/657f2acf01c70fb66fbcc9c06416123c7ed6.pdf
// This as reference: https://oeis.org/A003313
// Status: Complete
func problem122(n int) {
	children := make([]*AdditionTreeNode, 0)
	treeGenChain := []int{1}
	root := AdditionTreeNode{
		value:    1,
		children: children,
		parent:   nil,
	}

	depths := make(map[int]int)

	// Build Tree
	start := time.Now()
	buildAdditionTree(&root, treeGenChain, 0, depths)
	t := time.Now().Sub(start)
	fmt.Printf("Build tree time: %s\n", t.String())

	sum := 0
	for i := 1; i <= n; i++ {
		sum += depths[i]
	}

	fmt.Printf("%d", sum)
}

// There are some prime values, p, for which there exists a positive integer, n, such that the expression n^3 + p*n^2 is a perfect cube.
// Find how many primes below one million satisfy this property
// Status: Complete
func problem131(limit int) int {
	primes := helpers.FindPrimes(float64(limit))

	fmt.Println(len(primes))

	pMap := make(map[float64]bool)

	for _, p := range primes {
		pMap[float64(p)] = true
	}

	count := 0

	for k := 2; k < limit*limit; k++ {
		k3 := k * k * k
		for n := k - 1; n > 1; n-- {
			v := k3 - (n * n * n)
			if v%(n) == 0 {
				p := float64(v) / float64(n*n)
				_, ok := pMap[p]
				if ok {
					fmt.Printf("p=%d n=%d k=%d count=%d\n", int(p), n, int(k), count)
					count++
				}
			}
		}
	}

	return count
}

func main() {
	i := problem52()
	fmt.Println(i)
}
