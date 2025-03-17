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

	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}

	// sort based on first element of the pairs
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0] < pairs[j][0]
	})

	res := 1

	for start := n - 1; start >= 0; start-- {
		for end := start + 1; end < n; end++ {
			if pairs[end][0] <= pairs[start][1] {
				continue
			}
			dp[start] = max(dp[start], 1+dp[end])
			res = max(res, dp[start])
		}
	}

	return res
}
