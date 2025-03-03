package main

import "fmt"

func main() {
	matrix := [][]int{
		{2, 1, 3},
		{6, 5, 4},
		{7, 8, 9},
	}
	fmt.Println(minFallingPathSum(matrix))
}

const (
	MaxUint = ^uint(0)
	MaxInt  = int(MaxUint >> 1)
	MinInt  = -MaxInt - 1
)

func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)

	cache := make([][]int, n)
	for i := range cache {
		cache[i] = make([]int, n)
		for j := range cache[i] {
			cache[i][j] = MinInt
		}
	}

	var dfs func(i, j int) int

	dfs = func(i, j int) int {
		if i == n {
			return 0
		}
		if j >= n || j < 0 {
			return MaxInt
		}
		if cache[i][j] != MinInt {
			return cache[i][j]
		}
		res := matrix[i][j] + min(dfs(i+1, j-1), min(dfs(i+1, j), dfs(i+1, j+1)))
		cache[i][j] = res
		return res
	}

	res := MaxInt
	for i := 0; i < n; i++ {
		res = min(res, dfs(0, i))
	}

	return res
}
