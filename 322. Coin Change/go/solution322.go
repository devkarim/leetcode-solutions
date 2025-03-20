package main

import "fmt"

func main() {
	coins := []int{1, 2, 5}
	amount := 11
	fmt.Println(coinChange(coins, amount))
}

const (
	MaxUint = ^uint(0)
	MinUint = 0
	MaxInt  = int(MaxUint >> 1)
	MinInt  = -MaxInt - 1
)

func coinChange(coins []int, amount int) int {
	cache := make(map[int]int)

	var dfs func(total int) int

	dfs = func(total int) int {
		if total == 0 {
			return 0
		}
		if total < 0 {
			return MaxInt
		}
		if val, ok := cache[total]; ok {
			return val
		}
		res := MaxInt
		for _, coin := range coins {
			if coin > total {
				continue
			}
			take := dfs(total - coin)
			if take != MaxInt {
				res = min(res, take+1)
			}
		}
		cache[total] = res
		return res
	}

	res := dfs(amount)
	if res == MaxInt {
		return -1
	}

	return res
}
