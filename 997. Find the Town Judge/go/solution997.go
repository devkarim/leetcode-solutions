package main

import "fmt"

func main() {
	n := 3
	trust := [][]int{
		{1, 3}, // 1 -> 3
		{2, 3}, // 2 -> 3
	}
	fmt.Println(findJudge(n, trust))
}

func findJudge(n int, trust [][]int) int {
	ingoing := make(map[int]int)
	outgoing := make(map[int]int)
	for _, relationship := range trust {
		a := relationship[0]
		b := relationship[1]
		ingoing[b]++
		outgoing[a]++
	}
	for i := 1; i <= n; i++ {
		if ingoing[i] == n-1 && outgoing[i] == 0 {
			return i
		}
	}
	return -1
}
