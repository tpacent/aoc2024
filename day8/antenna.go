package day8

import (
	"aoc24/lib"
	"iter"
)

func Permute(items [][2]int) iter.Seq[[2][2]int] {
	return func(yield func([2][2]int) bool) {
		for k := 0; k < len(items)-1; k++ {
			for p := k + 1; p < len(items); p++ {
				if ok := yield([2][2]int{items[k], items[p]}); !ok {
					return
				}
			}
		}
	}
}

func Mirrors(a, b [2]int) [2][2]int {
	if a[0] > b[0] {
		a, b = b, a
	}

	dx := lib.AbsDiff(a[0], b[0])
	dy := lib.AbsDiff(a[1], b[1])
	xMin := min(a[0], b[0]) - dx
	xMax := max(a[0], b[0]) + dx
	yMin := min(a[1], b[1]) - dy
	yMax := max(a[1], b[1]) + dy

	if a[1] < b[1] {
		return [2][2]int{{xMin, yMin}, {xMax, yMax}}
	}

	return [2][2]int{{xMin, yMax}, {xMax, yMin}}
}

func PropagateSignal(a, b [2]int, steps, w, h int) iter.Seq[[2]int] {
	if a[0] > b[0] {
		a, b = b, a
	}

	resonant := steps < 0
	dx := b[0] - a[0]
	dy := b[1] - a[1]

	return func(yield func([2]int) bool) {
		if resonant {
			for _, coord := range [][2]int{a, b} {
				if ok := yield(coord); !ok {
					return
				}
			}
		}

		current := b

		for k := steps; k != 0; k-- {
			current[0] += dx
			current[1] += dy

			if !checkBounds(current, w, h) {
				break
			}

			if ok := yield(current); !ok {
				return
			}
		}

		current = a

		for k := steps; k != 0; k-- {
			current[0] -= dx
			current[1] -= dy

			if !checkBounds(current, w, h) {
				break
			}

			if ok := yield(current); !ok {
				return
			}
		}
	}
}

func checkBounds(coord [2]int, w, h int) bool {
	return coord[0] >= 0 &&
		coord[1] >= 0 &&
		coord[0] < w &&
		coord[1] < h
}
