package aoc1

func absDiffUint(a int, b int) int {
	if a < b {
		return b - a
	} else {
		return a - b
	}
}