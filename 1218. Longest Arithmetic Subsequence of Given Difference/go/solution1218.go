package main

import "fmt"

func main() {
	arr := []int{1, 5, 7, 8, 5, 3, 4, 2, 1}
	difference := -2
	fmt.Println(longestSubsequence(arr, difference))
}

func longestSubsequence(arr []int, difference int) int {
	dp := make(map[int]int)

	res := 1

	for _, num := range arr {
		if val, ok := dp[num-difference]; ok {
			dp[num] = val + 1
		} else {
			dp[num] = 1
		}
		res = max(res, dp[num])
	}

	return res
}
