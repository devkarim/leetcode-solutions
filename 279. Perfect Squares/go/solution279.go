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
	dp := make([]int, n+1)

	for i := range dp {
		dp[i] = n
	}
	dp[0] = 0

	for target := 1; target <= n; target++ {
		for s := 1; s <= m; s++ {
			square := s * s
			if target < square {
				break
			}
			dp[target] = min(dp[target], 1+dp[target-square])
		}
	}

	return dp[n]
}
