package day1_test

import (
	"aoc24/day1"
	"aoc24/lib"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestPartOne(t *testing.T) {
	l, r, err := prepareLists("testdata/one.txt")
	if err != nil {
		t.Fatal(err)
	}
	actual := day1.DiffLists(l, r)
	t.Log(actual) // 2176849
}

func TestPartTwo(t *testing.T) {
	l, r, err := prepareLists("testdata/one.txt")
	if err != nil {
		t.Fatal(err)
	}
	actual := day1.DiffMap(l, r)
	t.Log(actual) // 23384288
}

func prepareLists(filename string) (left, right []int, _ error) {
	input, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer input.Close()
	for vals := range lib.ReadInput(input, parseNums) {
		left = append(left, vals[0])
		right = append(right, vals[1])
	}
	return
}

func parseNums(s string) [2]int {
	flds := strings.Fields(s)
	a, _ := strconv.Atoi(flds[0])
	b, _ := strconv.Atoi(flds[1])
	return [2]int{a, b}
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
