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

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	res := 0

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if matrix[i][j] == '1' {
				down := dp[i+1][j]
				diag := dp[i+1][j+1]
				right := dp[i][j+1]
				dp[i][j] = 1 + min(down, diag, right)
				res = max(res, dp[i][j])
			}
		}
	}

	return res * res
}
