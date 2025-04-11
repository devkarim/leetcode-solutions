package main

import (
	"fmt"
	"math"
)

func main() {
	n := 13
	fmt.Println(numSquares(n))
}

func numSquares(n int) int {
	m := int(math.Sqrt(float64(n)))
	cache := make([][]int, n+1)

	for i := range cache {
		cache[i] = make([]int, m+1)
	}

	var dfs func(k, i int) int

	dfs = func(k, i int) int {
		if i > int(math.Sqrt(float64(k))) {
			if k == 0 {
				return 0
			}
			return n
		}
		if cache[k][i] != 0 {
			return cache[k][i]
		}
		// take
		take := 1 + dfs(k-i*i, i)
		// skip
		skip := dfs(k, i+1)
		cache[k][i] = min(take, skip)
		return cache[k][i]
	}

	return dfs(n, 1)
}
