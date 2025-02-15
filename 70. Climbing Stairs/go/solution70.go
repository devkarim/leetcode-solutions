package main

import "fmt"

func main() {
	n := 12
	res := climbStairs(n)
	fmt.Println(res)
}

/*
Problem Description:
Given N=n, how many possible ways to sum up to n?

Constraints:
- 1 <= n <= 45 -> relatively small
- Take either 1 step or 2 steps ahead

Example walkthrough:
- Given N=3, how many possible ways to up to 3?

-> Starting at N=3, what are my actions?

	-> Either take 1 step
	  -> Sum up to 2 (N-1)
	-> Either take 2 steps
	  -> Sum up to 1 (N-2)

Generalizing solution:
-> Starting at N=n, what are my actions?

	-> Either take 1 step
	  -> Sum up to (n-1)
	-> Either take 2 steps
	  -> Sum up to (n-2)

Base cases:
1) if n is either 0 or 1, we can climb only 1 step ahead
2) if n is 2, we can climb 1 step ahead twice or climb two steps at once

Recurrence relation:
f(n) = f(n-1) + f(n-2)
*/

var cache = make(map[int]int)

func climbStairs(n int) int {
	return dfs(n)
}

func dfs(n int) int {
	// if n is either 0 or 1, total number of ways to climb is 1
	if n <= 1 {
		return 1
	}
	// if n is 2, total number of ways to climb is 2 (1+1 or 2)
	if n == 2 {
		return 2
	}
	// is the result for n existing in cache?
	if val, ok := cache[n]; ok {
		return val
	}
	res := dfs(n-1) + dfs(n-2)
	cache[n] = res
	return cache[n]
}
