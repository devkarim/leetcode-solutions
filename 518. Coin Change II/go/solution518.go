package main

import "fmt"

func main() {
	amount := 5
	coins := []int{1, 2, 5}
	fmt.Println(change(amount, coins))
}

func change(amount int, coins []int) int {
	n := len(coins)
	cache := make([][]int, n)

	for i := range cache {
		cache[i] = make([]int, amount+1)
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}

	var dfs func(i, total int) int

	dfs = func(i, total int) int {
		if total == 0 {
			return 1
		}
		if i >= n || total < 0 {
			return 0
		}
		if cache[i][total] != -1 {
			return cache[i][total]
		}
		take := dfs(i, total-coins[i])
		skip := dfs(i+1, total)
		cache[i][total] = take + skip
		return cache[i][total]
	}

	return dfs(0, amount)
}
