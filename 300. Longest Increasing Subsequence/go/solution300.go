package main

import "fmt"

func main() {
	nums := []int{
		10,
		9,
		2,
		5,
		3,
		4,
		101,
		18,
	}

	fmt.Println(lengthOfLIS(nums))
}

func lengthOfLIS(nums []int) int {
	n := len(nums)

	cache := make([][]int, n)
	for i := range cache {
		cache[i] = make([]int, n)
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}

	var dfs func(curr, last int) int

	dfs = func(curr, last int) int {
		if curr >= n {
			return 0
		}
		if cache[curr][last+1] != -1 {
			return cache[curr][last+1]
		}
		// exclude current element
		res := dfs(curr+1, last)
		// include current element if it is the first element or
		// current item is larger than last item
		if last == -1 || nums[last] < nums[curr] {
			res = max(res, 1+dfs(curr+1, curr))
		}
		cache[curr][last+1] = res
		return res
	}

	return dfs(0, -1)
}
