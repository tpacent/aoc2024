package day6_test

import (
	"aoc24/day6"
	"aoc24/lib"
	"testing"
)

func TestInput(t *testing.T) {
	file := lib.MustOpenFile("testdata/input.txt")
	t.Cleanup(func() { _ = file.Close() })
	grid := lib.NewGrid(lib.ReadGrid(file, makeCell))
	walker := day6.NewWalker(grid, day6.State{X: 96, Y: 41, Dir: day6.DirU})
	t.Log(walker.Walk()) // 5067

}

func makeCell(b byte) day6.Cell {
	return day6.Cell{
		Empty:   b != '#',
		Visited: false,
	}
}
