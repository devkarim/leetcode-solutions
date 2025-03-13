package main

import "fmt"

func main() {
	s1 := "sea"
	s2 := "eat"
	fmt.Println(minimumDeleteSum(s1, s2))
}

func minimumDeleteSum(s1 string, s2 string) int {
	m := len(s1)
	n := len(s2)

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	countRemaining := func(s string, start int) int {
		cnt := 0
		for i := start; i < len(s); i++ {
			cnt = cnt + int(s[i])
		}
		return cnt
	}

	for i := 0; i < m; i++ {
		dp[i][n] = countRemaining(s1, i)
	}

	for j := 0; j < n; j++ {
		dp[m][j] = countRemaining(s2, j)
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if s1[i] == s2[j] {
				dp[i][j] = dp[i+1][j+1]
			} else {
				dp[i][j] = min(int(s1[i])+dp[i+1][j], int(s2[j])+dp[i][j+1])
			}
		}
	}

	return dp[0][0]
}
