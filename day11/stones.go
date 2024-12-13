package day11

import "aoc24/lib"

type Stone int

type StoneRule func(Stone) ([]Stone, bool)

func ApplyRules(stones []Stone, rules []StoneRule) (result []Stone) {
	for _, stone := range stones {
		for _, rule := range rules {
			if out, ok := rule(stone); ok {
				result = append(result, out...)
				break
			}
		}
	}
	return
}

func ApplyRuleSteps(stone Stone, rules []StoneRule, steps int) []Stone {
	result := []Stone{stone}
	for range steps {
		result = ApplyRules(result, rules)
	}
	return result
}

type CountState struct {
	TotalsCache map[int]map[Stone]int
	JumpCache   map[Stone][]Stone
	Rules       []StoneRule
	Stride      int
}

func CountStones(stones []Stone, state *CountState, level int) (total int) {
	if level == 0 {
		return len(stones)
	}

	if _, ok := state.TotalsCache[level]; !ok {
		state.TotalsCache[level] = make(map[Stone]int)
	}

	var nextBatch []Stone

	for _, stone := range stones {
		if value, ok := state.TotalsCache[level][stone]; ok {
			total += value
			continue
		}

		if value, ok := state.JumpCache[stone]; ok {
			nextBatch = value
		} else {
			nextBatch = ApplyRuleSteps(stone, state.Rules, state.Stride)
			state.JumpCache[stone] = nextBatch
		}

		value := CountStones(nextBatch, state, level-state.Stride)
		state.TotalsCache[level][stone] = value
		total += value
	}

	return
}

func RuleZero(stone Stone) ([]Stone, bool) {
	if stone != 0 {
		return nil, false
	}

	return []Stone{1}, true
}

func RuleEven(stone Stone) ([]Stone, bool) {
	digits := lib.NumDigits(int(stone))

	if digits%2 != 0 {
		return nil, false
	}

	left := stone

	mul := 1
	for range digits / 2 {
		left /= 10
		mul *= 10
	}

	right := stone - left*Stone(mul)
	return []Stone{left, right}, true
}

func RuleDefault(stone Stone) ([]Stone, bool) {
	return []Stone{2024 * stone}, true
}
