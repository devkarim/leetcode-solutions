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

/*
Problem Description:
Using an array of ints, find the maximum number of points by deleting an array to be earned
and then deleting element+1 and element-1

Observations:
# You are given an integer array nums. You want to maximize the number of points you get by
performing the following operation any number of times #
- Unordered array.
- Unlimited use of operation.
- Positive integers onlys => points always increasing
- Large input size
- Duplicates are allowed

# Pick any nums[i] and delete it to earn nums[i] points. #
- By picking a number, we can pick all the duplicates. Perhaps make a map for counts.
- Forced to delete the item picked to earn points.
- Array is always shrinking

# Afterwards, you must delete every element equal to nums[i] - 1 and every element equal to nums[i] + 1. #
- Forced to delete nums[i]-1 and nums[i] + 1 => delete every occurrence of them i.e. set count to zero.
- Maybe order array to properly remove duplicates?

# Return the maximum number of points you can earn by applying the above operation some number of times. #
- Optimal solution problem (max).
- Earning means taking the result of the operation.

Example Walkthrough:
Given [2,2,3,3,4,3], find the maximum number of earned points
- We can remove the duplicates by adding to hash map => [2,4,3]
- We can then order the array to easily remove nums[i+1] and nums[i-1] => [2,3,4]
- Now my array is [2,3,4] and hash map: (2=2,3=3,4=1)
-> What are my next actions?
- Take 2 (i=0) -> total=2*2=4
-- Cannot take 3 as it is nums[i]+1
-- Take 4 (i=2) -> total=4+4=8
- Donot take 2
-- Take 3 (i=1) -> total=3*3=9
-- Cannot take 4 as it is nums[i]+1

Generalizing:
=> Assuming we start at the end of the nums array:
- If next item is (n-1)
-- Take n (at position i) and move to i-2 => F(N) = nums[N] + F(N-2)
-- OR
-- Donot take n and move to i-1 => F(N) = F(N-1)
- If next item is not (n-1)
-- Take n (at position i) and move to i-1 => F(N) = nums[N] + F(N-1)
-- OR
-- Donot take n and move to i-1 => F(N) = F(N-1)

Base Cases:
- F(0) = 0
*/
func deleteAndEarn(nums []int) int {
	cnt := make(map[int]int)

	for _, k := range nums {
		cnt[k]++
	}

	uniqueNums := slices.Collect(maps.Keys(cnt))
	sort.Ints(uniqueNums)
	n := len(uniqueNums)

	dp := make([]int, n+2)

	dfs := func(idx, bias int) int {
		if idx < 0 {
			return 0
		}

		k := uniqueNums[idx]
		count := cnt[k]
		val := k * count
		res := 0

		if idx-1 >= 0 && uniqueNums[idx-1] == k-1 {
			res = max(val+dp[idx-2+bias], dp[idx-1+bias])
		} else {
			res = max(val+dp[idx-1+bias], dp[idx-1+bias])
		}

		return res
	}

	dp[0] = 0
	dp[1] = 0

	for i := 0; i < n; i++ {
		dp[i+2] = dfs(i, 2)
	}

	return dp[n+1]
}
