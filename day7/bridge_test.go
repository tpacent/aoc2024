package day7_test

import (
	"aoc24/day7"
	"aoc24/lib"
	"strings"
	"testing"
)

func TestPartOne(t *testing.T) {
	file := lib.MustOpenFile("testdata/input.txt")
	t.Cleanup(func() { _ = file.Close() })
	actual := 0

	ops := []day7.Op{day7.OpAdd, day7.OpMul}
	for task := range lib.ReadInput(file, parseLine) {
		if day7.Valid(task.Result, ops, task.Nums...) {
			actual += task.Result
		}
	}

	t.Log(actual) // 2941973819040
}

func TestPartTwo(t *testing.T) {
	file := lib.MustOpenFile("testdata/input.txt")
	t.Cleanup(func() { _ = file.Close() })
	actual := 0

	ops := []day7.Op{day7.OpAdd, day7.OpMul, day7.OpCat}
	for task := range lib.ReadInput(file, parseLine) {
		if day7.Valid(task.Result, ops, task.Nums...) {
			actual += task.Result
		}
	}

	t.Log(actual) // 249943041417600
}

const exampledata = `
190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

type Task struct {
	Result int
	Nums   []int
}

func TestExample(t *testing.T) {
	actual := 0
	ops := []day7.Op{day7.OpAdd, day7.OpMul}
	for task := range lib.ReadInput(strings.NewReader(strings.TrimSpace(exampledata)), parseLine) {
		if day7.Valid(task.Result, ops, task.Nums...) {
			actual += task.Result
		}
	}
	if actual != 3749 {
		t.Error("unexpected", actual)
	}
}

func TestExample2(t *testing.T) {
	actual := 0
	ops := []day7.Op{day7.OpAdd, day7.OpMul, day7.OpCat}
	for task := range lib.ReadInput(strings.NewReader(strings.TrimSpace(exampledata)), parseLine) {
		if day7.Valid(task.Result, ops, task.Nums...) {
			actual += task.Result
		}
	}
	if actual != 11387 {
		t.Error("unexpected", actual)
	}
}

func parseLine(s string) (task Task) {
	result, numsLine, _ := strings.Cut(s, ":")
	task.Result = lib.MustParse(result)
	for _, value := range strings.Fields(numsLine) {
		task.Nums = append(task.Nums, lib.MustParse(value))
	}
	return
}
