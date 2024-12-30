package day18

import "aoc24/lib"

type Finder struct {
	grid    *lib.Grid[byte]
	visited *lib.Grid[int]
	result  int
}

func NewFinder(g *lib.Grid[byte]) *Finder {
	return &Finder{
		grid: g,
	}
}

func (f *Finder) Walk(from, to lib.Coords, pathExists bool) int {
	f.result = 0
	w, h := f.grid.Width(), f.grid.Height()
	f.visited = lib.NewGrid(w, h, make([]int, w*h))
	f.walk(from, to, 0, pathExists)
	return f.result
}

var deltas = []lib.Coords{
	{X: 0, Y: -1},
	{X: 1, Y: 0},
	{X: 0, Y: 1},
	{X: -1, Y: 0},
}

func (f *Finder) walk(loc, end lib.Coords, step int, pathExists bool) {
	if loc == end {
		if f.result == 0 {
			f.result = step
		} else {
			f.result = min(f.result, step)
		}

		return
	}

	if pathExists && f.result > 0 {
		return
	}

	prev, didVisit := f.visited.At(loc.X, loc.Y)

	if didVisit && prev > 0 && prev <= step {
		return // dead end
	}

	f.visited.Set(loc.X, loc.Y, step)

	for _, delta := range deltas {
		next := lib.Coords{X: loc.X + delta.X, Y: loc.Y + delta.Y}

		if tile, ok := f.grid.At(next.X, next.Y); !ok || tile == '#' {
			continue // wall
		}

		f.walk(next, end, step+1, pathExists)
	}
}
