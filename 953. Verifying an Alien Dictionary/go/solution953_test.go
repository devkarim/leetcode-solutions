package main

import "testing"

func Test_isAlienSorted(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		words []string
		order string
		want  bool
	}{
		{
			name: "Test Case #1",
			words: []string{
				"hello",
				"leetcode",
			},
			order: "hlabcdefgijkmnopqrstuvwxyz",
			want:  true,
		},
		{
			name: "Test Case #2",
			words: []string{
				"word",
				"world",
				"row",
			},
			order: "worldabcefghijkmnpqstuvxyz",
			want:  false,
		},
		{
			name: "Test Case #3",
			words: []string{
				"apple",
				"app",
			},
			order: "abcdefghijklmnopqrstuvwxyz",
			want:  false,
		},
		{
			name: "Test Case #4",
			words: []string{
				"ubg",
				"kwh",
			},
			order: "qcipyamwvdjtesbghlorufnkzx",
			want:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isAlienSorted(tt.words, tt.order)
			if got != tt.want {
				t.Errorf("isAlienSorted() = %v, want %v", got, tt.want)
			}
		})
	}
}
