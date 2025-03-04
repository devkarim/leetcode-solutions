package main

import "fmt"

func main() {
	matrix := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	fmt.Println(maximalSquare(matrix))
}

func maximalSquare(matrix [][]byte) int {
	m := len(matrix)
	n := len(matrix[0])

	cache := make([][]int, m)
	for i := range cache {
		cache[i] = make([]int, n)
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}

	var dfs func(i, j int) int

	dfs = func(i, j int) int {
		if i >= m || j >= n {
			return 0
		}
		if cache[i][j] != -1 {
			return cache[i][j]
		}
		up := dfs(i+1, j)
		diag := dfs(i+1, j+1)
		left := dfs(i, j+1)
		cache[i][j] = 0
		if matrix[i][j] == '1' {
			res := 1 + min(up, diag, left)
			cache[i][j] = res
		}
		return cache[i][j]
	}

	dfs(0, 0)

	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res = max(res, cache[i][j])
		}
	}

	return res * res
}
