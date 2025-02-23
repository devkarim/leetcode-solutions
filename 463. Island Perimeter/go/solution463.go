package main

import "fmt"

func main() {
	grid := [][]int{
		{0, 1, 0, 0},
		{1, 1, 1, 0},
		{0, 1, 0, 0},
		{1, 1, 0, 0},
	}
	fmt.Println(islandPerimeter(grid))
}

func islandPerimeter(grid [][]int) int {
	rows := len(grid)    // number of rows
	cols := len(grid[0]) // number of cols

	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	var dfs func(int, int) int

	dfs = func(r, c int) int {
		if r >= rows || r < 0 || c >= cols || c < 0 || grid[r][c] == 0 {
			return 1
		}
		if visited[r][c] {
			return 0
		}
		visited[r][c] = true
		return dfs(r, c+1) + dfs(r, c-1) + dfs(r+1, c) + dfs(r-1, c)
	}

	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == 1 {
				return dfs(r, c)
			}
		}
	}

	return -1
}
