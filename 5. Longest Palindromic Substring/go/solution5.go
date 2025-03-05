package main

import "fmt"

func main() {
	s := "babad"
	fmt.Println(longestPalindrome(s))
}

func longestPalindrome(s string) string {
	n := len(s)

	expandFromCenter := func(l, r int) string {
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
		}
		return s[l+1 : r]
	}

	res := ""

	for i := 0; i < n; i++ {
		odd := expandFromCenter(i, i)
		even := expandFromCenter(i, i+1)

		if len(odd) > len(res) {
			res = odd
		}

		if len(even) > len(res) {
			res = even
		}
	}

	return res
}
