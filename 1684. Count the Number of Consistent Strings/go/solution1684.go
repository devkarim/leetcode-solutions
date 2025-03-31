package main

import (
	"fmt"
	"strings"
)

func main() {
	allowed := "ab"
	words := []string{"ad", "bd", "aaab", "baa", "badab"}
	result := countConsistentStrings(allowed, words)
	fmt.Println(result) // Output: 2
}

func countConsistentStrings(allowed string, words []string) int {
	res := 0

	isConsistent := func(word string) bool {
		for _, c := range word {
			if !strings.ContainsRune(allowed, c) {
				return false
			}
		}
		return true
	}

	for _, w := range words {
		if isConsistent(w) {
			res++
		}
	}

	return res
}
