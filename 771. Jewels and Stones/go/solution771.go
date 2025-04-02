package main

func main() {
	jewels := "aA"
	stones := "aAAbbbb"
	result := numJewelsInStones(jewels, stones)
	println(result) // Output: 3
}

func numJewelsInStones(jewels string, stones string) int {
	isJewel := func(r rune) bool {
		for _, j := range jewels {
			if j == r {
				return true
			}
		}
		return false
	}

	res := 0

	for _, s := range stones {
		if isJewel(s) {
			res++
		}
	}

	return res
}
