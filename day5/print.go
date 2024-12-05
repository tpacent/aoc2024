package day5

import "slices"

func NewPrintQueue(rules map[int][]int) *PrintQueue {
	return &PrintQueue{
		precedence: rules,
	}
}

type PrintQueue struct {
	precedence map[int][]int
}

func (pq *PrintQueue) Validate(update []int) bool {
	pageIndices := make(map[int]int, len(update))

	for index, id := range update {
		pageIndices[id] = index
	}

	for index, id := range update {
		for _, before := range pq.precedence[id] {
			beforeIndex, ok := pageIndices[before]

			if !ok {
				continue
			}

			if index > beforeIndex {
				return false
			}
		}
	}

	return true
}

func (pq *PrintQueue) Fix(update []int) ([]int, bool) {
	if pq.Validate(update) {
		return update, false
	}

	slices.SortFunc(update, func(a, b int) int {
		if slices.Contains(pq.precedence[a], b) {
			return -1
		}

		return 1
	})

	return update, true
}

func Tuples2Rules(tuples [][2]int) map[int][]int {
	rules := map[int][]int{}
	for _, tuple := range tuples {
		before, after := tuple[0], tuple[1]
		afterRules := rules[before]
		rules[before] = append(afterRules, after)
	}
	return rules
}
