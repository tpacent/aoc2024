package day11_test

import (
	"aoc24/day11"
	"aoc24/lib"
	"bytes"
	"os"
	"testing"
)

func TestPartOne(t *testing.T) {
	input, err := os.ReadFile("testdata/input.txt")
	if err != nil {
		t.Fatal(err)
	}

	state := &day11.CountState{
		TotalsCache: make(map[int]map[day11.Stone]int),
		JumpCache:   make(map[day11.Stone][]day11.Stone),
		Rules:       rules,
		Stride:      5,
	}

	t.Log(day11.CountStones(parseInput(input), state, 25)) // 184927
}

func TestPartTwo(t *testing.T) {
	input, err := os.ReadFile("testdata/input.txt")
	if err != nil {
		t.Fatal(err)
	}

	state := &day11.CountState{
		TotalsCache: make(map[int]map[day11.Stone]int),
		JumpCache:   make(map[day11.Stone][]day11.Stone),
		Rules:       rules,
		Stride:      5,
	}

	t.Log(day11.CountStones(parseInput(input), state, 75)) // 220357186726677
}

var rules = []day11.StoneRule{
	day11.RuleZero,
	day11.RuleEven,
	day11.RuleDefault,
}

func TestExample(t *testing.T) {
	stones := []day11.Stone{125, 17}

	state := &day11.CountState{
		TotalsCache: make(map[int]map[day11.Stone]int),
		JumpCache:   make(map[day11.Stone][]day11.Stone),
		Rules:       rules,
		Stride:      1,
	}

	if actual := day11.CountStones(stones, state, 25); actual != 55312 {
		t.Error("unexpected", actual)
	}
}

func parseInput(input []byte) (stones []day11.Stone) {
	for _, s := range bytes.Fields(input) {
		stones = append(stones, day11.Stone(lib.MustParse(string(s))))
	}
	return
}
