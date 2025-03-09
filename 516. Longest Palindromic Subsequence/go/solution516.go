package main

import "fmt"

func main() {
	s := "bbbab"
	fmt.Println(longestPalindromeSubseq(s))
}

func longestPalindromeSubseq(s string) int {
	n := len(s)
	r := reverse(s)

	dp := make([][]int, n+1)

	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := n - 1; i <= 0; i-- {
		for j := n - 1; j <= 0; j-- {
			if s[i] == r[j] {
				dp[i][j] = 1 + dp[i+1][j+1]
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}

	return dp[0][0]
}

func reverse(s string) string {
	reversed := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		reversed[i], reversed[j] = reversed[j], reversed[i]
	}
	return string(reversed)
}
