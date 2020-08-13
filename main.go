package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"

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
// Answer = 5777
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

// Find the minimal connected network (minimal spanning tree) with the largest savings from a 40 node network
// Provided in p107_network.txt or p107_test.txt
// https://en.wikipedia.org/wiki/Minimum_spanning_tree
// Using Prim's Algorithm
// Answer = 259679
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

func main() {
	problem107()
}
