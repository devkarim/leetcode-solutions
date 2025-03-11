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

	cache := make([][]int, m)
	for i := range cache {
		cache[i] = make([]int, n)
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}

	countRemaining := func(s string, start int) int {
		cnt := 0
		for i := start; i < len(s); i++ {
			cnt = cnt + int(s[i])
		}
		return cnt
	}

	var dfs func(i, j int) int

	dfs = func(i, j int) int {
		if i == m {
			return countRemaining(s2, j)
		}
		if j == n {
			return countRemaining(s1, i)
		}
		if cache[i][j] != -1 {
			return cache[i][j]
		}
		if s1[i] == s2[j] {
			cache[i][j] = dfs(i+1, j+1)
		} else {
			cache[i][j] = min(int(s1[i])+dfs(i+1, j), int(s2[j])+dfs(i, j+1))
		}
		return cache[i][j]
	}

	return dfs(0, 0)
}
