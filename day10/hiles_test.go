package day10_test

import (
	"aoc24/day10"
	"aoc24/lib"
	"strings"
	"testing"
)

func TestPartOne(t *testing.T) {
	file := lib.MustOpenFile("testdata/input.txt")
	t.Cleanup(func() { _ = file.Close() })

	grid := lib.NewGrid(lib.ReadGrid(file, func(b byte) day10.TrailPos {
		return day10.TrailPos{Height: uint8(b - '0')}
	}))

	total := 0

	for thead := range grid.FindAll(pTrailHead) {
		total += len(day10.WalkTrail(thead, grid.Clone()))
	}

	lib.PrintResult(t, 10, 1, total, 667)
}

func TestPartTwo(t *testing.T) {
	file := lib.MustOpenFile("testdata/input.txt")
	t.Cleanup(func() { _ = file.Close() })

	grid := lib.NewGrid(lib.ReadGrid(file, func(b byte) day10.TrailPos {
		return day10.TrailPos{Height: uint8(b - '0')}
	}))

	rating := 0
	for thead := range grid.FindAll(pTrailHead) {
		for tail := range day10.WalkTrail(thead, grid.Clone()) {
			rating += day10.TrailRating(thead, lib.WithCoords[day10.TrailPos]{Coords: tail}, grid.Clone())
		}
	}

	lib.PrintResult(t, 10, 2, rating, 1344)
}

func pTrailHead(wc lib.WithCoords[day10.TrailPos]) bool {
	return wc.Value.Height == 0
}

const example = `
..90..9
...1.98
...2..7
6543456
765.987
876....
987....
`

func TestExample(t *testing.T) {
	grid := lib.NewGrid(lib.ReadGrid(strings.NewReader(example), func(b byte) day10.TrailPos {
		return day10.TrailPos{Height: uint8(b - '0')}
	}))

	thead, ok := grid.Find(pTrailHead)

	if !ok {
		t.Fatal("trailhead exists")
	}

	if nines := day10.WalkTrail(thead, grid.Clone()); len(nines) != 4 {
		t.Error("unexpected trails count", len(nines))
	}
}

const example2 = `
012345
123456
234567
345678
4.6789
56789.
`

func TestExample2(t *testing.T) {
	grid := lib.NewGrid(lib.ReadGrid(strings.NewReader(example2), func(b byte) day10.TrailPos {
		return day10.TrailPos{Height: uint8(b - '0')}
	}))

	thead, ok := grid.Find(pTrailHead)

	if !ok {
		t.Fatal("trailhead exists")
	}

	rating := 0
	for coords := range day10.WalkTrail(thead, grid.Clone()) {
		rating += day10.TrailRating(thead, lib.WithCoords[day10.TrailPos]{Coords: coords}, grid.Clone())
	}

	if rating != 227 {
		t.Error("unexpected rating", rating)
	}
}
