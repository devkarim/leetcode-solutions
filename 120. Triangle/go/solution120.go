package main

import "fmt"

func main() {
	triangle := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}
	fmt.Println(minimumTotal(triangle))
}

const (
	MaxUint = ^uint(0)
	MinUint = 0
	MaxInt  = int(MaxUint >> 1)
)

func minimumTotal(triangle [][]int) int {
	m := len(triangle)
	n := len(triangle[m-1])

	dp := make([]int, n)
	copy(dp, triangle[m-1])

	for i := m - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			dp[j] = triangle[i][j] + min(dp[j+1], dp[j])
		}
	}

	return dp[0]
}
