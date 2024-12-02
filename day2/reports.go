package day2

import "aoc24/lib"

const safeDiffMax = 3
const safeDiffMin = 1

func IsSafe(levels []int) bool {
	var predicate func(a, b int) bool

	if levels[0] < levels[1] {
		predicate = isInc
	} else {
		predicate = isDec
	}

	for k := 1; k < len(levels); k++ {
		a, b := levels[k-1], levels[k]

		if !predicate(a, b) {
			return false
		}

		if lib.AbsDiff(a, b) > safeDiffMax {
			return false
		}
	}

	return true
}

func IsSafeRecover(levels []int) bool {
	if IsSafe(levels) {
		return true
	}

	for k := 0; k < len(levels); k++ {
		if IsSafe(Pluck(levels, k)) {
			return true
		}
	}

	return false
}

func IsSafeDamped(levels []int, errors int) bool {
	if errors < 0 {
		return false
	}

	var predicate func(a, b int) bool

	if levels[0] < levels[1] {
		predicate = isInc
	} else {
		predicate = isDec
	}

	for k := 1; k < len(levels); k++ {
		a, b := levels[k-1], levels[k]

		if !predicate(a, b) {
			return IsSafeDamped(Pluck(levels, k-1), errors-1)
		}

		if lib.AbsDiff(a, b) > safeDiffMax || lib.AbsDiff(a, b) < safeDiffMin {
			return IsSafeDamped(Pluck(levels, k), errors-1) || IsSafeDamped(Pluck(levels, k-1), errors-1)
		}
	}

	return true
}

func Pluck(levels []int, index int) []int {
	out := make([]int, len(levels)-1)
	copy(out, levels[:index])
	copy(out[index:], levels[index+1:])
	return out
}

func isInc(a, b int) bool {
	return a < b
}

func isDec(a, b int) bool {
	return a > b
}

func Violations(levels []int) int {
	var (
		incs, decs int
		diffs      int
	)

	for k := 1; k < len(levels); k++ {
		a, b := levels[k-1], levels[k]

		if !safeDiff(a, b) {
			diffs++
		}

		if a > b {
			decs++
		}

		if a < b {
			incs++
		}
	}

	return diffs + min(incs, decs)
}

func safeDiff(a, b int) bool {
	diff := lib.AbsDiff(a, b)

	if diff < safeDiffMin {
		return false
	}

	if diff > safeDiffMax {
		return false
	}

	return true
}
