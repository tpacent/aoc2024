package day1_test

import (
	"aoc24/day1"
	"aoc24/lib"
	"testing"
)

func TestPartOne(t *testing.T) {
	actual := day1.DiffLists(prepareLists("testdata/input.txt"))
	lib.PrintResult(t, 1, 1, actual, 2176849)
}

func TestPartTwo(t *testing.T) {
	actual := day1.DiffMap(prepareLists("testdata/input.txt"))
	lib.PrintResult(t, 1, 2, actual, 23384288)
}

func prepareLists(filename string) (left, right []int) {
	file := lib.MustOpenFile(filename)
	defer file.Close()
	for vals := range lib.ReadInput(file, lib.NumsLine) {
		left = append(left, vals[0])
		right = append(right, vals[1])
	}
	return
}

func TestExample(t *testing.T) {
	const expected = 11

	actual := day1.DiffLists(
		[]int{3, 4, 2, 1, 3, 3},
		[]int{4, 3, 5, 3, 9, 3},
	)

	if actual != expected {
		t.Fatal("unexpected value")
	}
}
