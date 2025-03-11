package main

import "fmt"

func main() {
	word1 := "horse"
	word2 := "ros"
	fmt.Println(minDistance(word1, word2))
}

/*
Problem Description
Using two strings, word1 and word2. Find the minimum number of operations to transform word1 to word2.
The operations are:
- Insert a character
- Replace a character
- Delete a character

Observations
- Low input size.
- Lowercase English letters are allowed only.
- Counting problem (minimum i.e. optimal solution is needed).
- We are allowed to either insert, replace or delete a character from word1 ONLY.
- Given word1="" and word2="", the result is: 0.
- Given word1="abc" and word2="", the result is: 3 (delete "abc").
- Given word1="" and word2="abc", the result is: 3 (insert "abc").
- Given word1="a" and word2="b", the result is: 1 (replace "a" with "b").
- Given word1="abc" and word2="abc", the result is: 0 (no change is needed).

Example Walkthrough
Given word1="abc" and word2="adc". Find the minimum number of operations.

Assume i is the index for word1 and j is the index for word2.
Starting at i=0 and j=0, both characters are equal (a=a) so move to next character in both strings.
- Move to i=1 (i+1) and j=1 (j+1)
-- Both characters do not match (b!=d) so we have three choices:
-- Insert: (i, j+1) -> Insert jth character at position i so we need to advance j only.
-- Replace: (i+1,j+1) -> Replace character at i with character at j (i.e words[i]==word[j] at this moment) so we need to advance both, i and j.
-- Delete: (i+1,j) -> Delete the character at i hoping the rest of word2 equals word1 after that so we need to advance i only
*/
func minDistance(word1, word2 string) int {
	n := len(word1)
	m := len(word2)

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	for i := 0; i < n; i++ {
		dp[i][m] = n - i
	}

	for j := 0; j < m; j++ {
		dp[n][j] = m - j
	}

	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if word1[i] == word2[j] {
				dp[i][j] = dp[i+1][j+1]
			} else {
				dp[i][j] = 1 + min(dp[i+1][j+1], dp[i+1][j], dp[i][j+1])
			}
		}
	}

	return dp[0][0]
}
