package main

import "fmt"

func main() {
	text1 := "abcde"
	text2 := "ace"
	fmt.Println(longestCommonSubsequence(text1, text2))
}

func longestCommonSubsequence(text1 string, text2 string) int {
	n := len(text1)
	m := len(text2)

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if text1[i] == text2[j] {
				dp[i][j] = 1 + dp[i+1][j+1]
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}

	return dp[0][0]
}
