package main

import "fmt"

func main() {
	m := 3
	n := 2
	fmt.Println(uniquePaths(m, n))
}

/*
Problem Description:
Using a size of grid (m,n) where m represents the number of rows and n represents the number of columns.
Count the number of unique paths starting from coordinates (0,0) to (m-1,n-1)

Observations:
# There is a robot on an m x n grid. The robot is initially located at the top-left corner (i.e., grid[0][0]). #
- Forced to start at (0,0).
- Small input size.
- Robot means the moving entity.
- Given grid 1x1 => count = 1.
- Given grid 1x2 => count = 1.
- Given grid 2x2 => count = 2.
# The robot tries to move to the bottom-right corner (i.e., grid[m - 1][n - 1]). #
- Forced to end at (m-1,n-1).
# The robot can only move either down or right at any point in time. #
- Can either move right or down only.
- Going down means (i+1, j) if the current position is (i,j).
- Going right means (i, j+1) if the current position is (i,j).
- If out of bounds, we aren't at the final destination (i.e. return zero).
# Given the two integers m and n, return the number of possible unique paths that the robot can take to
reach the bottom-right corner. #
- Counting problem.
- Counting only unique paths, increasing always.
- A path means going to bottom-right corner successfully starting at top-left corner.
- A unique path is a path that is not discovered before.
# The test cases are generated so that the answer will be less than or equal to 2 * 10^9. #
- Answer is always less than 2x10^9
- Integer is sufficient to hold the count of unique paths.
- Size of integer is 4 bytes => (2^32)/2 roughly equals 2.15x10^9 data range.

Example Walkthrough
Given m=rows=3, n=cols=2. Find the unique paths from top-left to bottom-right using
bottom and right directions only.

=> Starting at (0,0), we have two actions:
- Move right (0,1)
-- Move right (0,2)
-- Move bottom (1,1)
- Move bottom (1,0)
-- Move right (1,1)
-- Move bottom (2,0)

Generalizing
Assuming we are in (i,j). We have two actions:
- Move right (i, j+1)
-- Move right (i,j+2)
-- Move bottom (i+1,j+1) -> (1)
- Move bottom (i+1,j)
-- Move right (i+1,j+1) -> (2)
-- Move bottom (i+2,j)

From (1) and (2), there is overlap i.e. calculationg are made more than once.
Therefore, we can cache the result in a 2D array of size (m,n) same as grid size.

Recurrence Relation
F(i,j) = F(i,j+1) + F(i+1,j)

Base Cases
Assuming a grid of size (m,n)
-> F(m-1,n-1) = 1 => At destination
-> Let x >= m, y >= n
--> F(x,j)=0 => Moved to right and out of bounds
--> F(i,y)=0 => Moved to bottom and out of bounds
*/
func uniquePaths(m, n int) int {
	dp := make([][]int, m+1)

	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i+1][j] + dp[i][j+1]
			}
		}
	}

	return dp[0][0]
}
