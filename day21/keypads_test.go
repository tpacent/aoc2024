package day21_test

import (
	"aoc24/day21"
	"aoc24/lib"
	"bufio"
	"io"
	"testing"
)

func TestDay21Part1(t *testing.T) {
	file := lib.MustOpenFile("testdata/input.txt")
	t.Cleanup(func() { _ = file.Close() })

	actual := 0

	for _, code := range ParseInput(file) {
		actual += day21.CodeComplexity([]byte(code), 2)
	}

	t.Log(actual) // 270084
}

func TestDay21Part2(t *testing.T) {
	file := lib.MustOpenFile("testdata/input.txt")
	t.Cleanup(func() { _ = file.Close() })

	actual := 0

	for _, code := range ParseInput(file) {
		actual += day21.CodeComplexity([]byte(code), 25)
	}

	t.Log(actual) // 329431019997766
}

func TestExample(t *testing.T) {
	actual := 0

	for _, code := range []string{"029A", "980A", "179A", "456A", "379A"} {
		actual += day21.CodeComplexity([]byte(code), 2)
	}

	if actual != 126384 {
		t.Error("unexpected value")
	}
}

func ParseInput(r io.Reader) (codes []string) {
	for scanner := bufio.NewScanner(r); scanner.Scan(); {
		if line := scanner.Text(); len(line) > 0 {
			codes = append(codes, line)
		}
	}
	return
}
