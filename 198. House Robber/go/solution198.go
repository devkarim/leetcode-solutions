package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(rob(nums))
}

/*
Problem Description:
Given a houses array, where ith house has nums[i] money. Find the
max amount of money that can be robbed without alerting the police.

Constraints:
- 1 <= nums.length <= 100 -> small
- 0 <= nums[i] <= 400 -> small
- Cannot take adjacent values in houses array
- Forced to start at either first house or second house

Observations:
# Each house has a certain amount of money stashed, the only constraint
stopping you from robbing each of them is that adjacent houses have security
systems connected and it will automatically contact the police if two adjacent
houses were broken into on the same night. #
- Each item in the array indicates a house money.
- Array contains only positive and zero.
-- Stealing either increases total money or not at all.
- Contains duplicate elements which makes sense as each house is an independent entity.
- Cannot take adjacent elements at once.
- Order of array matters to check adjacency.

# Given an integer array nums representing the amount of money of each house,
return the maximum amount of money you can rob tonight without alerting the police. #
- Optimal solution problem (maximum).
- Alerting the police means no adjacent houses are robbed at once.
- Robbing means making a choice to take ith house money i.e. ith array element.

Example Walkthrough:
Given [2,1,1,2], find max of either starting at 1st element or 2nd array element with
ignoring adjacent elements.

-> What are my actions?
- Take 1st house
-- Take 3rd house
-- OR
-- Take 4th house
- OR
- Take 2nd house
-- Take 4th house

Conclusion:
-> If we start at 1st house, either take 3rd or 4th.
-> If we start at 2nd house, either take 4th or 5th.

Generalizing:
Assuming we start at house: i
-> Take house money of: i
-> Get max of (i + 2) and (i + 3)

Recurrence Relation:
-> F(N) = nums[N] + Max(F(N+2), F(N+3)) <-

Base Cases:
-> F(N) = 0 at: N >= len(nums)
*/
var cache []int

func rob(nums []int) int {
	cache = make([]int, len(nums))
	for i := range cache {
		cache[i] = -1
	}
	return max(dfs(nums, 0), dfs(nums, 1))
}

func dfs(nums []int, idx int) int {
	if idx >= len(nums) {
		return 0
	}
	if cache[idx] != -1 {
		return cache[idx]
	}
	res := nums[idx] + max(dfs(nums, idx+2), dfs(nums, idx+3))
	cache[idx] = res
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
