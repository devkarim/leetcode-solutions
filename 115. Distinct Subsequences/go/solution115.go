package main

import "fmt"

func main() {
	s := "babgbag"
	t := "bag"
	fmt.Println(numDistinct(s, t))
}

func numDistinct(s string, t string) int {
	m := len(s)
	n := len(t)

	cache := make([][]int, m)
	for i := range cache {
		cache[i] = make([]int, n)
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}

	var dfs func(i, j int) int

	dfs = func(i, j int) int {
		// if reached the end of t then we found a match
		if j >= n {
			return 1
		}
		// if reached the end of s then we didn't find a match
		if i >= m {
			return 0
		}
		if cache[i][j] != -1 {
			return cache[i][j]
		}
		// if both characters are equal, advance j and i OR just advance i to try other subsequences
		if s[i] == t[j] {
			cache[i][j] = dfs(i+1, j+1) + dfs(i+1, j)
		} else {
			// if both characters are not equal, advance only i since we didn't find a matching character
			cache[i][j] = dfs(i+1, j)
		}
		return cache[i][j]
	}

	return dfs(0, 0)
}
