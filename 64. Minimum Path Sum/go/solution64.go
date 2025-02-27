package main

import (
	"fmt"
)

func main() {
	grid := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	// grid := [][]int{
	// 	{1, 2, 3},
	// 	{4, 5, 6},
	// }
	fmt.Println(minPathSum(grid))
}

const (
	MaxUint = ^uint(0)
	MinUint = 0
	MaxInt  = int(MaxUint >> 1)
	MinInt  = -MaxInt - 1
)

/*
Problem Description
Using a grid of size m x n, find the minimum cost to reach bottom-right corner starting at top-left corner.

Observations
- Contains positive numbers ond zero only.
- Small input size.
- Forced to start at (0,0).
- Forced to end at (m-1,n-1).
- Optimal solution problem (minimum).
- Forced to move either bottom or right only.
- If at position (i,j), moving right means (i,j+1)
- If at position (i,j), moving bottom means (i+1,j)

Example Walkthrough
Given [[1,3,1],[1,5,1],[4,2,1]], find the minimum cost to reach the bottom-right corner.

=> Starting at (0,0), we have two actions and cost=1:
- Move right (0,1) => Cost=1+3=4
-- Move right (0,2) => Cost=4+1=5
-- Move bottom (1,1) => Cost=4+5=9
- Move bottom (1,0) => Cost=1+1=2
-- Move right (1,1) => Cost=2+5=7
-- Move bottom (2,0) => Cost=2+4=6

Generalizing
Assuming we start at (i,j). We have two actions plus take current cost at (i,j) (i.e. cost=grid[i][j]):
- Move right (i, j+1) => Cost+=grid[i][j+1]
-- Move right (i,j+2) => Cost+=grid[i][j+2]
-- Move bottom (i+1,j+1) -> (1) => Cost+=grid[i+1][j+1]
- Move bottom (i+1,j) => Cost+=grid[i][j+2]
-- Move right (i+1,j+1) -> (2) => Cost+=grid[i+1][j+1]
-- Move bottom (i+2,j) => Cost+=grid[i+2][j]

From (1) and (2), there is overlap i.e. calculations are made more than once for the same input.
Therefore, we can cache the result in a 2D array of size (m,n) same as grid size.

Recurrence Relation
F(i,j) = grid[i][j] + min(F(i+1,j), F(i,j+1))

Base Cases
Assuming a grid of size (m,n)
-> F(m-1,n-1) = grid[i][j] => At destination
-> Let x >= m, y >= n
--> F(x,j)=Infinity => Moved to right and out of bounds
--> F(i,y)=Infinity => Moved to bottom and out of bounds
*/
func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	cache := make([][]int, m)
	for i := range cache {
		cache[i] = make([]int, n)
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}

	var dfs func(i, j int) int

	dfs = func(i, j int) int {
		// if out of bounds
		if i >= m || j >= n {
			return MaxInt
		}
		// if at destination
		if i == m-1 && j == n-1 {
			return grid[i][j]
		}
		// if in cache
		if cache[i][j] != -1 {
			return cache[i][j]
		}
		// move either right or left
		res := grid[i][j] + min(dfs(i+1, j), dfs(i, j+1))
		cache[i][j] = res
		return res
	}

	return dfs(0, 0)
}
