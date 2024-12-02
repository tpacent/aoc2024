package day1

import (
	"slices"
)

func DiffLists(left, right []int) (total int) {
	slices.Sort(left)
	slices.Sort(right)
	for k := 0; k < len(left); k++ {
		total += diff(left[k], right[k])
	}
	return
}

func DiffMap(left, right []int) (total int) {
	counts := make(map[int]int)

	for _, n := range right {
		counts[n]++
	}

	for _, n := range left {
		total += n * counts[n]
	}

	return
}

func diff(a, b int) int {
	if a > b {
		return a - b
	}

	return b - a
}
