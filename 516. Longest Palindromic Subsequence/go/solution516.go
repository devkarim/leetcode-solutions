package main

import "fmt"

func main() {
	s := "bbbab"
	fmt.Println(longestPalindromeSubseq(s))
}

func longestPalindromeSubseq(s string) int {
	n := len(s)
	r := reverse(s)

	cache := make([][]int, n)

	for i := range cache {
		cache[i] = make([]int, n)
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}

	var dfs func(i, j int) int

	dfs = func(i, j int) int {
		if i < 0 || j < 0 {
			return 0
		}
		if cache[i][j] != -1 {
			return cache[i][j]
		}
		res := 0
		if s[i] == r[j] {
			res = 1 + dfs(i-1, j-1)
		} else {
			res = max(dfs(i-1, j), dfs(i, j-1))
		}
		cache[i][j] = res
		return res
	}

	return dfs(n-1, n-1)
}

func reverse(s string) string {
	reversed := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		reversed[i], reversed[j] = reversed[j], reversed[i]
	}
	return string(reversed)
}
