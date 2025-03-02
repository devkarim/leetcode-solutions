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

func minimumTotal(triangle [][]int) int {
	m := len(triangle)
	n := len(triangle[m-1])

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	dfs := func(i, j int) int {
		return triangle[i][j] + min(dp[i+1][j+1], dp[i+1][j])
	}

	for i := m - 1; i >= 0; i-- {
		for j := len(triangle[i]) - 1; j >= 0; j-- {
			dp[i][j] = dfs(i, j)
		}
	}

	return dp[0][0]
}
