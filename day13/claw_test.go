package day13_test

import (
	"aoc24/day13"
	"aoc24/lib"
	"bufio"
	"io"
	"iter"
	"regexp"
	"testing"
)

func TestPartOne(t *testing.T) {
	file := lib.MustOpenFile("testdata/input.txt")
	t.Cleanup(func() { _ = file.Close() })

	actual := day13.CalcTokens(parseInput(file), 0)
	lib.PrintResult(t, 13, 1, actual, 31589)
}

func TestPartTwo(t *testing.T) {
	file := lib.MustOpenFile("testdata/input.txt")
	t.Cleanup(func() { _ = file.Close() })

	actual := day13.CalcTokens(parseInput(file), 10000000000000)
	lib.PrintResult(t, 13, 2, actual, 98080815200063)
}

func parseInput(src io.Reader) iter.Seq[day13.Input] {
	scanner := bufio.NewScanner(src)

	return func(yield func(day13.Input) bool) {
		var input day13.Input
		var counter int

		for scanner.Scan() {
			counter++
			line := scanner.Text()

			if line == "" {
				counter = 0
				continue
			}

			nums := reNumbers.FindAllString(line, 2)

			x := lib.MustParse(nums[0])
			y := lib.MustParse(nums[1])

			switch counter {
			case 1:
				input.AButton = day13.Button{Dx: x, Dy: y}
			case 2:
				input.BButton = day13.Button{Dx: x, Dy: y}
			case 3:
				input.Target = day13.Target{X: x, Y: y}
				if ok := yield(input); !ok {
					return
				}
			}
		}

	}
}

var reNumbers = regexp.MustCompile(`\d+`)
