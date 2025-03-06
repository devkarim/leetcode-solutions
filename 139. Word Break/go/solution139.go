package main

import "fmt"

func main() {
	s := "leetcode"
	wordDict := []string{
		"leet",
		"code",
	}
	fmt.Println(wordBreak(s, wordDict))
}

/*
Problem Description:
Using a string and a dictonary of words, return true if the string can be space separated using the dictionary of words.

Observations:
- All words in the wordDict array are unique. Maybe use hashset?
- The same word can be used multiple times in the string "s".
- Low input size.
- Unordered array.
- Only lowercase English letters are in the array and the string.

Solution Approach
- For each word in wordDict, check whether the word can fit the string provided.
- If it can fit:
-- Let k be the length of the chosen word.
-- If substring of "s" at a starting position "i" ending at "k" equals the chosen word:
--- Shift i to be i+k, the new starting position then repeat the above again.
--- If ended at the last of the string, then it is certainly a valid word.
-- Otherwise, continue to next word
*/
func wordBreak(s string, wordDict []string) bool {
	n := len(s)

	dp := make([]bool, n+1)
	dp[n] = true

	for start := n - 1; start >= 0; start-- {
		for _, w := range wordDict {
			end := start + len(w)
			if end <= n && s[start:end] == w {
				dp[start] = dp[end]
			}
			if dp[start] {
				break
			}
		}
	}

	return dp[0]
}
