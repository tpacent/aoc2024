package day20_test

import (
	"aoc24/day20"
	"aoc24/lib"
	"bufio"
	"io"
	"strings"
	"testing"
)

func TestPartOne(t *testing.T) {
	file := lib.MustOpenFile("testdata/input.txt")
	t.Cleanup(func() { _ = file.Close() })
	track, start, end := parseInput(file)
	solver := day20.NewSolver(track)
	solver.Prepare(start, end)
	lib.PrintResult(t, 20, 1, solver.CountCheats(2, 100), 1311)
}

func TestPartTwo(t *testing.T) {
	file := lib.MustOpenFile("testdata/input.txt")
	t.Cleanup(func() { _ = file.Close() })
	track, start, end := parseInput(file)
	solver := day20.NewSolver(track)
	solver.Prepare(start, end)
	lib.PrintResult(t, 20, 2, solver.CountCheats(20, 100), 961364)
}

const example = `
###############
#...#...#.....#
#.#.#.#.#.###.#
#S#...#.#.#...#
#######.#.#.###
#######.#.#...#
#######.#.###.#
###..E#...#...#
###.#######.###
#...###...#...#
#.#####.#.###.#
#.#...#.#.#...#
#.#.#.#.#.#.###
#...#...#...###
###############
`

func TestExample(t *testing.T) {
	tcases := []struct {
		Name     string
		Radius   int
		Thresh   int
		Expected int
	}{
		{Name: "one", Radius: 2, Thresh: 2, Expected: 44},
		{Name: "two", Radius: 20, Thresh: 50, Expected: 285},
	}

	track, start, end := parseInput(strings.NewReader(example))
	solver := day20.NewSolver(track)
	solver.Prepare(start, end)

	for _, tt := range tcases {
		t.Run(tt.Name, func(t *testing.T) {
			if actual := solver.CountCheats(tt.Radius, tt.Thresh); actual != tt.Expected {
				t.Errorf("unexpected value %d; expected %d", actual, tt.Expected)
			}
		})
	}
}

func parseInput(r io.Reader) (_ *lib.Grid[byte], start lib.Coords, end lib.Coords) {
	scanner := bufio.NewScanner(r)
	var data []byte
	var w, h int

	for scanner.Scan() {
		line := scanner.Bytes()

		if len(line) == 0 {
			continue
		}

		for x, c := range line {
			switch c {
			case 'S':
				start = lib.Coords{X: x, Y: h}
				line[x] = '.'
			case 'E':
				end = lib.Coords{X: x, Y: h}
				line[x] = '.'
			}
		}

		data = append(data, line...)
		w = len(line)
		h++
	}

	return lib.NewGrid(w, h, data), start, end
}
