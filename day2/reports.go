package day2

import "aoc24/lib"

const safeDiffMax = 3
const safeDiffMin = 1

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
			for _, pos := range []int{k - 2, k - 1, k} {
				if pos >= 0 && IsSafeDamped(Pluck(levels, pos), errors-1) {
					return true
				}
			}

			return false
		}

		if !safeDiff(a, b) {
			return IsSafeDamped(Pluck(levels, k-1), errors-1) || IsSafeDamped(Pluck(levels, k), errors-1)
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
