package main

import "fmt"

func main() {
	nums1 := []int{1, 4, 2}
	nums2 := []int{1, 2, 4}
	fmt.Println(maxUncrossedLines(nums1, nums2))
}

func maxUncrossedLines(nums1 []int, nums2 []int) int {
	m := len(nums1)
	n := len(nums2)

	cache := make([][]int, m)
	for i := range cache {
		cache[i] = make([]int, n)
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}

	var dfs func(i, j int) int
	
	dfs = func(i, j int) int {
		if i == m || j == n {
			return 0
		}
		if cache[i][j] != -1 {
			return cache[i][j]
		}
		if nums1[i] == nums2[j] {
			cache[i][j] = dfs(i+1, j+1)+1
		} else {
			cache[i][j] = max(dfs(i+1, j), dfs(i, j+1))
		}
		return cache[i][j]
	}

	return dfs(0, 0)
}
