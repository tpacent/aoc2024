package day3_test

import (
	"aoc24/day3"
	"aoc24/lib"
	"os"
	"testing"
)

func TestPartOne(t *testing.T) {
	data, err := os.ReadFile("testdata/input.txt")
	if err != nil {
		t.Fatal(err)
	}

	actual := day3.MulSum(string(data))
	lib.PrintResult(t, 3, 1, actual, 187194524)
}

func TestPartTwo(t *testing.T) {
	data, err := os.ReadFile("testdata/input.txt")
	if err != nil {
		t.Fatal(err)
	}

	actual := day3.MulSumToggle(string(data))
	lib.PrintResult(t, 3, 2, actual, 127092535)
}

func TestExample(t *testing.T) {
	const in = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"

	if day3.MulSum(in) != 161 {
		t.Error("unexpected")
	}
}

func TestExample2(t *testing.T) {
	const in = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"

	if day3.MulSumToggle(in) != 48 {
		t.Error("unexpected")
	}
}
