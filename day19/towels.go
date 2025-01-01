package day19

func CountArrangements(design string, patterns map[string]struct{}) int {
	return countArrangements(design, patterns, map[uint32]int{}, 0)
}

func countArrangements(design string, patterns map[string]struct{}, cache map[uint32]int, index int) (count int) {
	if len(design) == index {
		return 1
	}

	if value, ok := cache[uint32(index)]; ok {
		return value
	}

	for k := 1; k < len(design)-index+1; k++ {
		if _, ok := patterns[design[index:index+k]]; ok {
			count += countArrangements(design, patterns, cache, index+k)
		}
	}

	cache[uint32(index)] = count
	return
}
