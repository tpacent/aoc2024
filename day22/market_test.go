package day22_test

import (
	"aoc24/day22"
	"aoc24/lib"
	"slices"
	"testing"
)

func TestDay22Part1(t *testing.T) {
	input := lib.MustOpenFile("testdata/input.txt")

	iter := lib.ReadInput(input, func(s string) uint32 {
		return uint32(lib.MustParse(s))
	})

	total := 0
	for n := range iter {
		total += int(day22.NthStep(n, 2000))
	}

	t.Log(total) // 20401393616
}

func TestDay22Part2(t *testing.T) {
	input := lib.MustOpenFile("testdata/input.txt")

	secrets := slices.Collect(lib.ReadInput(input, func(s string) uint32 {
		return uint32(lib.MustParse(s))
	}))

	pattern, value := day22.MaxPattern(secrets, 2000+1) // off by one

	t.Log(value, pattern) // 2272
}

func TestSteps(t *testing.T) {
	var n uint32 = 123

	for range 10 {
		n = day22.Step(n)
	}

	if actual := day22.NthStep(123, 10); 5908254 != actual {
		t.Error("unexpected value")
	}
}

func TestCalcStocks(t *testing.T) {
	stocks := day22.CalcStocks(123, 10)
	if value := stocks.Patterns[[4]int8{-1, -1, 0, 2}]; value != 6 {
		t.Error("unexpected value")
	}
}

func TestExample(t *testing.T) {
	pattern, value := day22.MaxPattern([]uint32{1, 2, 3, 2024}, 2000)

	if pattern != [4]int8{-2, 1, -1, 3} || value != 23 {
		t.Log("unexpected value")
	}
}
