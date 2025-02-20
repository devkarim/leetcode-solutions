package main

import "fmt"

func main() {
	nums := []int{2, 3, 2}
	fmt.Println(rob(nums))
}

var cache []int

func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	cache = make([]int, n)
	fill(cache, -1)
	res := dfs(nums, n-2, 0)
	fill(cache, -1)
	res = max(res, dfs(nums, n-1, 1))
	return res
}

func dfs(nums []int, idx int, last int) int {
	if idx < last {
		return 0
	}
	if cache[idx] != -1 {
		return cache[idx]
	}
	res := max(nums[idx]+dfs(nums, idx-2, last), dfs(nums, idx-1, last))
	cache[idx] = res
	return res
}

func fill(arr []int, val int) {
	for i := range arr {
		arr[i] = val
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
