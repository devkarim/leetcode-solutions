package main

import (
	"fmt"
	"strings"
)

func main() {
	words := []string{"ubg", "kwh"}
	order := "qcipyamwvdjtesbghlorufnkzx"
	fmt.Println(isAlienSorted(words, order))
}

/*
Problem Description:
Using an array and order, check if the array is sorted lexicographically based on order.

Observations:
- Low input size.
- All characters are lower case.
- Each word has a maximum of 20 characters.
- Order is exactly 26 characters.
- Order has unique characters, maybe use hash map?
- In English, "ab" comes before "abc".
- In English, "a" comes before "b".

Example Walkthrough:
Given words=["word", "world", "row"], order="worldabcefghijkmnpqstuvxyz" check if this array is sorted.

-> At col=0: w = w < r -> Ordered
-> At col=1: o=o=o -> Ordered
-> At col=2: r=r -> Ordered
-> At col=3: d > l -> Unordered

Generalizing:
- Look for the first different character between two words, say A and B (i.e. [A,B]).
-- If the first different character is found:
--- Let c1 be A's character and c2 be B's character
--- If c1 > c2 => Not ordered
-- If the first different character is not found:
--- If B doesn't start with A (i.e. B doesn't have a prefix of A) => Not ordered
*/
func isAlienSorted(words []string, order string) bool {
	n := len(words)
	if n == 1 {
		return true
	}

	orderMap := make(map[rune]int)
	for i, w := range order {
		orderMap[w] = i
	}

	for i, a := range words {
		if i == n-1 {
			break
		}
		b := words[i+1]
		firstDiffRuneAt := firstDiffRune(a, b)
		if firstDiffRuneAt != -1 {
			c1 := a[firstDiffRuneAt]
			c2 := b[firstDiffRuneAt]
			if orderMap[rune(c1)] > orderMap[rune(c2)] {
				return false
			}
		} else {
			if !strings.HasPrefix(b, a) {
				return false
			}
		}
	}

	return true
}

func firstDiffRune(a, b string) int {
	shortest := a
	if len(shortest) > len(b) {
		shortest = b
	}
	for i := range shortest {
		if a[i] != b[i] {
			return i
		}
	}
	return -1
}
