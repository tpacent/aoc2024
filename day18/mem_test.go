package day18_test

import (
	"bufio"
	"io"
	"strings"
	"testing"

	"aoc24/day18"
	"aoc24/lib"
)

func TestPartOne(t *testing.T) {
	grid, _ := prepareGrid(t)
	finder := day18.NewFinder(grid)
	steps := finder.Walk(lib.Coords{}, lib.Coords{X: 70, Y: 70}, false)
	lib.PrintResult(t, 18, 1, steps, 298)
}

func TestPartTwo(t *testing.T) {
	grid, coords := prepareGrid(t)
	finder := day18.NewFinder(grid)
	var result lib.Coords

	for k := 1024; k < len(coords); k++ {
		c := coords[k]
		grid.Set(c.X, c.Y, '#')

		if 0 == finder.Walk(lib.Coords{}, lib.Coords{X: 70, Y: 70}, true) {
			result = c
			break
		}

	}
	lib.PrintResult(t, 18, 2, lib.JoinInts([]int{result.X, result.Y}), "52,32")
}

func prepareGrid(t *testing.T) (*lib.Grid[byte], []lib.Coords) {
	r := lib.MustOpenFile("testdata/input.txt")
	t.Cleanup(func() { _ = r.Close() })
	coords := parseInput(r)
	grid := lib.NewGrid(71, 71, make([]byte, 71*71))
	for _, c := range coords[:1024] {
		grid.Set(c.X, c.Y, '#')
	}
	return grid, coords
}

var example = `
5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0
`

func TestExample(t *testing.T) {
	grid := lib.NewGrid(7, 7, make([]byte, 49))

	for _, coord := range parseInput(strings.NewReader(example))[:12] {
		grid.Set(coord.X, coord.Y, '#')
	}

	finder := day18.NewFinder(grid)
	actual := finder.Walk(lib.Coords{X: 0, Y: 0}, lib.Coords{X: 6, Y: 6}, false)
	if actual != 22 {
		t.Error("unexpected value")
	}
}

func parseInput(r io.Reader) (coords []lib.Coords) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		l, r, _ := strings.Cut(line, ",")
		coords = append(coords, lib.Coords{X: lib.MustParse(l), Y: lib.MustParse(r)})
	}
	return
}
