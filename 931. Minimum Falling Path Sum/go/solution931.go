package main

import "fmt"

func main() {
	matrix := [][]int{
		{17, 82},
		{1, -44},
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
	if n == 1 {
		return matrix[0][0]
	}

	dp := make([]int, n)
	copy(dp, matrix[n-1][:])
	res := MaxInt

	for i := n - 2; i >= 0; i-- {
		left := MaxInt
		for j := 0; j < n; j++ {
			mid := dp[j]
			right := MaxInt
			if j+1 < n {
				right = dp[j+1]
			}
			dp[j] = matrix[i][j] + min(left, mid, right)
			left = mid
			if i == 0 {
				res = min(res, dp[j])
			}
		}
	}

	return res
}
