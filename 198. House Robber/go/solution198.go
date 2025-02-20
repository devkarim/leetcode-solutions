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
- Forced to start at either first house or second house or last house

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
Given [2,1,1,2], find max of either starting at the last element with ignoring
adjacent elements.

-> What are my actions?
- Rob last house (index = 3)
-- Take last house money = 2
-- Move to third last house (index = 1)
- OR
- Donot rob last house (index = 3)
-- Move to second last house (index = 2)

Conclusion:

Generalizing:
Assuming we start at house: i
-> Take house money of: i
--> Move to i - 2
-> Donot take house money of: i
--> Move to i - 1

Recurrence Relation:
-> F(N) = Max(nums[N]+F(N-2), F(N-1)) <-

Base Cases:
-> F(0) = 0
*/
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]int, len(nums))

	dp[0] = nums[0]
	dp[1] = max(nums[1], nums[0])

	for i := 2; i < len(nums); i++ {
		dp[i] = max(nums[i]+dp[i-2], dp[i-1])
	}

	return dp[len(nums)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
