package main

import (
	"fmt"
)

func main() {
	cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	// cost := []int{10, 15, 20}
	fmt.Println(minCostClimbingStairs(cost))
}

/*
Problem Description:
Given a cost array, each stair has a cost to pair for.
Calculate the minimum cost to reach the top by either taking
one step or two steps ahead.

Constraints:
- 2 <= cost.length <= 1000 -> small i/p size
- 0 <= cost[i] <= 999 -> small values
- Take either 1 step or steps

Observations:
# You are given an integer array cost where cost[i] is the cost of ith step on a staircase.
	Once you pay the cost, you can either climb one or two steps. #
- Each stair has an integer cost.
- Only positive costs are given, no negatives.
-- Advancing only increases the total cost unless zero is given.
- Can climb either one step or two steps ahead.
- Cannot take same stair twice (makes sense).
- Climbing means advancing one or two steps.
- Costs are not unique.

# You can either start from the step with index 0, or the step with index 1. #
- At the beginning, we can start at 1st stair or 2nd stair.
- Forced to start at specified position (0 or 1)

# Return the minimum cost to reach the top of the floor. #
- Optimal solution problem (minimum)
- Top means out of array boundaries => stopping condition.

Example Walkthrough
- Given cost=[10, 15, 20], find the minimum cost to reach the top.
-> What are my actions?
	-> Take first cost (start at first index)
	-> Total cost: 10
		-> Take second cost
		-> Total cost: 25
			-> Take third cost
			-> Total cost: 45
		-> Take third cost
		-> Total cost: 30
	-> Take second cost
	-> Total cost: 15 <-- Minimum solution

Conclusion:
We have two choices, either step one or two steps.
-> If we step one, take 1st cost
-> If we step twice, take 2nd cost

Generalizing:
-> Take ith step
-> Total cost: cost[i] => actions: i + 1, i + 2

-> Take ith + 1 step
-> Total cost: cost[i + 1] => actions: i + 2, i + 3

Recurrence Relation:
-> F(N) = Min(cost[N] + F(N+1), cost[N+1] + F(N+2)) <-

Base Cases:
-> F(len(costs)) = 0
*/

func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost)+2)
	dfs := func(cost []int, idx int) int {
		if idx >= len(cost) {
			return 0
		}
		// include first cost
		res := cost[idx] + dp[idx+1]
		// include second cost
		second := 0
		if idx+1 < len(cost) {
			second = cost[idx+1]
		}
		res = min(res, second+dp[idx+2])
		return res
	}
	for i := len(cost) - 1; i >= 0; i-- {
		dp[i] = dfs(cost, i)
	}
	return dp[0]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
