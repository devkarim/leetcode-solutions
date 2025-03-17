package main

import (
	"fmt"
	"sort"
)

func main() {
	pairs := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
	}
	fmt.Println(findLongestChain(pairs))
}

func findLongestChain(pairs [][]int) int {
	n := len(pairs)

	cache := make([][]int, n)
	for i := range cache {
		cache[i] = make([]int, n)
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}

	// sort based on first element of the pairs
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0] < pairs[j][0]
	})

	var dfs func(curr, last int) int

	dfs = func(curr, last int) int {
		if curr == n {
			return 0
		}
		if cache[curr][last+1] != -1 {
			return cache[curr][last+1]
		}
		// exclude curr element
		res := dfs(curr+1, last)
		if last == -1 || pairs[curr][0] > pairs[last][1] {
			// include curr element
			res = max(res, 1+dfs(curr+1, curr))
		}
		cache[curr][last+1] = res
		return res
	}

	return dfs(0, -1)
}
