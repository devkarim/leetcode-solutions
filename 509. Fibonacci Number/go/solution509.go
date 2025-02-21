package main

import "fmt"

func main() {
	n := 30
	fmt.Println(fib(n))
}

func fib(n int) int {
	if n <= 1 {
		return n
	}

	beforeBefore := 0
	before := 1

	for i := 2; i <= n; i++ {
		beforeBefore, before = before, before+beforeBefore
	}

	return before
}
