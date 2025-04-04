package day5_test

import (
	"aoc24/day5"
	"aoc24/lib"
	"bufio"
	"strings"
	"testing"
)

func TestPartOne(t *testing.T) {
	file := lib.MustOpenFile("testdata/input.txt")
	t.Cleanup(func() { _ = file.Close() })

	scanner := bufio.NewScanner(file)
	rules := day5.Tuples2Rules(parseTuples(scanner))
	queue := day5.NewPrintQueue(rules)
	actual := 0

	for scanner.Scan() {
		update := readUpdate(scanner.Text())
		if queue.Validate(update) {
			actual += update[len(update)/2]
		}
	}

	lib.PrintResult(t, 5, 1, actual, 3608)
}

func TestPartTwo(t *testing.T) {
	file := lib.MustOpenFile("testdata/input.txt")
	t.Cleanup(func() { _ = file.Close() })

	scanner := bufio.NewScanner(file)
	rules := day5.Tuples2Rules(parseTuples(scanner))
	queue := day5.NewPrintQueue(rules)
	actual := 0

	for scanner.Scan() {
		update := readUpdate(scanner.Text())
		if update, fixed := queue.Fix(update); fixed {
			actual += update[len(update)/2]
		}
	}

	lib.PrintResult(t, 5, 2, actual, 4922)
}

func readUpdate(line string) (update []int) {
	for c := range strings.SplitSeq(line, ",") {
		update = append(update, lib.MustParse(c))
	}
	return
}

func parseTuples(scanner *bufio.Scanner) (tuples [][2]int) {
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		l, r, _ := strings.Cut(scanner.Text(), "|")
		tuples = append(tuples, [2]int{lib.MustParse(l), lib.MustParse(r)})
	}

	return
}
