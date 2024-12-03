package day2_test

import (
	"aoc24/day2"
	"aoc24/lib"
	"fmt"
	"testing"
)

func TestPartOne(t *testing.T) {
	file := lib.MustOpenFile("testdata/input.txt")
	t.Cleanup(func() { _ = file.Close() })

	var safeReports int

	for r := range lib.ReadInput(file, lib.NumsLine) {
		if day2.IsSafeDamped(r, 0) {
			safeReports++
		}
	}

	t.Log(safeReports) // 218
}

func TestPartTwo(t *testing.T) {
	file := lib.MustOpenFile("testdata/input.txt")
	t.Cleanup(func() { _ = file.Close() })

	var safeReports int

	for r := range lib.ReadInput(file, lib.NumsLine) {
		if day2.IsSafeDamped(r, 1) {
			safeReports++
		}
	}

	t.Log(safeReports) // 290
}

func TestExample(t *testing.T) {
	tests := []struct {
		Report []int
		Expect bool
	}{
		{Report: []int{7, 6, 4, 2, 1}, Expect: true},
		{Report: []int{1, 2, 7, 8, 9}, Expect: false},
		{Report: []int{9, 7, 6, 2, 1}, Expect: false},
		{Report: []int{1, 3, 2, 4, 5}, Expect: false},
		{Report: []int{8, 6, 4, 4, 1}, Expect: false},
		{Report: []int{1, 3, 6, 7, 9}, Expect: true},
	}

	for index, tcase := range tests {
		t.Run(fmt.Sprintf("test %d", index), func(t *testing.T) {
			actual := day2.IsSafeDamped(tcase.Report, 0)
			if actual != tcase.Expect {
				t.Error("unexpected")
			}
		})
	}
}

func TestExample2(t *testing.T) {
	tests := []struct {
		Report []int
		Expect bool
	}{
		{Report: []int{7, 6, 4, 2, 1}, Expect: true},
		{Report: []int{1, 2, 7, 8, 9}, Expect: false},
		{Report: []int{9, 7, 6, 2, 1}, Expect: false},
		{Report: []int{1, 3, 2, 4, 5}, Expect: true},
		{Report: []int{8, 6, 4, 4, 1}, Expect: true},
		{Report: []int{1, 3, 6, 7, 9}, Expect: true},
	}

	for index, tcase := range tests {
		t.Run(fmt.Sprintf("test %d", index), func(t *testing.T) {
			actual := day2.IsSafeDamped(tcase.Report, 1)
			if actual != tcase.Expect {
				t.Error("unexpected")
			}
		})
	}
}
