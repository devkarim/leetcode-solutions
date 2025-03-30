package main

import "fmt"

func main() {
	pushed := []int{1, 2, 3, 4, 5}
	popped := []int{4, 5, 3, 2, 1}
	result := validateStackSequences(pushed, popped)
	fmt.Println(result) // Output: true

	pushed = []int{1, 2, 3, 4, 5}
	popped = []int{4, 3, 5, 1, 2}
	result = validateStackSequences(pushed, popped)
	fmt.Println(result) // Output: false
}

func validateStackSequences(pushed []int, popped []int) bool {
	var stack []int

	popPtr := 0

	for _, v := range pushed {
		stack = append(stack, v)
		for len(stack) != 0 && stack[len(stack)-1] == popped[popPtr] {
			stack = stack[:len(stack)-1]
			popPtr++
		}
	}

	return len(stack) == 0
}
