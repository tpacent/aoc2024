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

	t.Log(calcTotal(grid)) // 1533024
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

	if actual := calcTotal(grid); actual != 1930 {
		t.Error("unexpected", actual)
	}
}

func calcTotal(grid *lib.Grid[day12.Plot]) (total int) {
	for _, cell := range grid.Iter() {
		if cell.Value.Visited {
			continue
		}
		area, perimeter := day12.FloodFill(cell, grid)
		total += area * perimeter
	}
	return
}
