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
	cache := make([][]int, n)
	for i := range cache {
		cache[i] = make([]int, m)
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
		if text1[i] == text2[j] {
			res = 1 + dfs(i-1, j-1)
		} else {
			res = max(dfs(i-1, j), dfs(i, j-1))
		}
		cache[i][j] = res
		return res
	}

	return dfs(n-1, m-1)
}
