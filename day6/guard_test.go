package day6_test

import (
	"aoc24/day6"
	"aoc24/lib"
	"strings"
	"testing"
)

func TestPartOne(t *testing.T) {
	file := lib.MustOpenFile("testdata/input.txt")
	t.Cleanup(func() { _ = file.Close() })
	grid := lib.NewGrid(lib.ReadGrid(file, makeCell))
	init, _ := grid.Find(startpos)
	t.Log(day6.NewWalker(grid).Walk(init.X, init.Y, day6.DirU)) // 5067
}

func TestPartTwo(t *testing.T) {
	file := lib.MustOpenFile("testdata/input.txt")
	t.Cleanup(func() { _ = file.Close() })
	grid := lib.NewGrid(lib.ReadGrid(file, makeCell))
	init, _ := grid.Find(startpos)
	walker := day6.NewWalker(grid)
	walker.Walk(init.X, init.Y, day6.DirU)
	visited := walker.Visited()
	actual := 0
	for _, pos := range grid.Iter() {
		if _, ok := visited[[2]int{pos.X, pos.Y}]; ok {
			newgrid := grid.Clone()
			newgrid.Set(pos.X, pos.Y, '#')
			if _, ok := day6.NewWalker(newgrid).Walk(init.X, init.Y, day6.DirU); !ok {
				actual++
			}
		}
	}

	t.Log(actual) // 1793
}

var labmap = `
....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

func TestExample(t *testing.T) {
	grid := lib.NewGrid(lib.ReadGrid(strings.NewReader(labmap), makeCell))
	init, _ := grid.Find(startpos)
	actual, ok := day6.NewWalker(grid).Walk(init.X, init.Y, day6.DirU)
	if ok != true {
		t.Error("unexpected loop")
	}
	if actual != 41 {
		t.Log("unexpected value")
	}
}

func TestExample2(t *testing.T) {
	grid := lib.NewGrid(lib.ReadGrid(strings.NewReader(labmap), makeCell))
	init, _ := grid.Find(startpos)
	walker := day6.NewWalker(grid)
	walker.Walk(init.X, init.Y, day6.DirU) // fill visited map
	visited := walker.Visited()
	looped := 0
	for _, pos := range grid.Iter() {
		if _, ok := visited[[2]int{pos.X, pos.Y}]; ok {
			newgrid := grid.Clone()
			newgrid.Set(pos.X, pos.Y, '#')
			if _, ok := day6.NewWalker(newgrid).Walk(init.X, init.Y, day6.DirU); !ok {
				looped++
			}
		}
	}
	if looped != 6 {
		t.Error("unexpected value", looped)
	}
}

func makeCell(b byte) byte { return b }

func startpos(wc lib.WithCoords[byte]) bool { return wc.Value == '^' }
