package main

import (
	"fmt"
	"maps"
	"slices"
	"sort"
)

func main() {
	nums := []int{3, 4, 2}
	fmt.Println(deleteAndEarn(nums))
}

func deleteAndEarn(nums []int) int {
	cnt := make(map[int]int)

	for _, n := range nums {
		cnt[n]++
	}

	uniqueNums := slices.Collect(maps.Keys(cnt))
	sort.Ints(uniqueNums)
	n := len(uniqueNums)

	cache := make([]int, n)

	var dfs func(idx int) int

	dfs = func(idx int) int {
		if idx >= n {
			return 0
		}
		if cache[idx] != 0 {
			return cache[idx]
		}

		k := uniqueNums[idx]
		count := cnt[k]
		val := k * count
		res := 0

		// if next item is k+1
		if idx+1 < n && uniqueNums[idx+1] == k+1 {
			res = max(val+dfs(idx+2), dfs(idx+1))
		} else {
			res = max(val+dfs(idx+1), dfs(idx+1))
		}

		cache[idx] = res

		return res
	}

	return dfs(0)
}
