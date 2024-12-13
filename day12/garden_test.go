package day12_test

import (
	"aoc24/day12"
	"aoc24/lib"
	"strings"
	"testing"
)

func TestPartOne(t *testing.T) {
	file := lib.MustOpenFile("testdata/input.txt")
	t.Cleanup(func() { _ = file.Close() })
	grid := lib.NewGrid(lib.ReadGrid(file, func(b byte) day12.Plot {
		return day12.Plot{Type: b}
	}))

	t.Log(day12.CalcTotalPerimeter(grid)) // 1533024
}

func TestPartTwo(t *testing.T) {
	file := lib.MustOpenFile("testdata/input.txt")
	t.Cleanup(func() { _ = file.Close() })
	grid := lib.NewGrid(lib.ReadGrid(file, func(b byte) day12.Plot {
		return day12.Plot{Type: b}
	}))

	t.Log(day12.CalcTotalWalls(grid)) // 910066
}

const example = `
RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`

func TestExample(t *testing.T) {
	grid := lib.NewGrid(lib.ReadGrid(strings.NewReader(example), func(b byte) day12.Plot {
		return day12.Plot{Type: b}
	}))

	if actual := day12.CalcTotalPerimeter(grid); actual != 1930 {
		t.Error("unexpected", actual)
	}
}

const example2 = `
AAAAAA
AAABBA
AAABBA
ABBAAA
ABBAAA
AAAAAA`

func TestExample2(t *testing.T) {
	grid := lib.NewGrid(lib.ReadGrid(strings.NewReader(example2), func(b byte) day12.Plot {
		return day12.Plot{Type: b}
	}))

	if actual := day12.CalcTotalWalls(grid); actual != 368 {
		t.Error("unexpected", actual)
	}
}
