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

	cache := make([][]int, m)
	for i := range cache {
		cache[i] = make([]int, n)
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}

	var dfs func(i, j int) int

	dfs = func(i, j int) int {
		if i >= m || j >= i+1 {
			return 0
		}
		if cache[i][j] != -1 {
			return cache[i][j]
		}
		cache[i][j] = triangle[i][j] + min(dfs(i+1, j+1), dfs(i+1, j))
		return cache[i][j]
	}

	return dfs(0, 0)
}
