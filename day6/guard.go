package day6

import "aoc24/lib"

type Dir uint8

const (
	DirUnknown = 0
	DirU       = 0b0001
	DirD       = 0b0010
	DirL       = 0b0100
	DirR       = 0b1000
)

var dirs = map[Dir][2]int{
	DirU: {0, -1},
	DirD: {0, 1},
	DirL: {-1, 0},
	DirR: {1, 0},
}

func NewWalker(grid *lib.Grid[byte]) *Walker {
	return &Walker{
		grid:    grid,
		visited: make(map[[2]int]Dir),
	}
}

type Walker struct {
	grid    *lib.Grid[byte]
	visited map[[2]int]Dir
}

func (w *Walker) Walk(x, y int, dir Dir) (int, bool) {
	for {
		if dirs, ok := w.visited[[2]int{x, y}]; ok && (dirs&dir > 0) {
			return len(w.visited), false
		} else {
			w.visited[[2]int{x, y}] |= dir
		}

		dirDeltas := dirs[dir]
		dx, dy := dirDeltas[0], dirDeltas[1]
		nx, ny := x+dx, y+dy
		nextCell, ok := w.grid.At(nx, ny)

		if !ok {
			break // outside grid
		}

		if nextCell == '#' { // true = blocked
			dir = nextDir(dir)
			continue
		}

		x = nx
		y = ny
	}

	return len(w.visited), true
}

func (w *Walker) Visited() map[[2]int]Dir {
	return w.visited
}

func nextDir(dir Dir) Dir {
	switch dir {
	case DirU:
		return DirR
	case DirR:
		return DirD
	case DirD:
		return DirL
	case DirL:
		return DirU
	}

	panic("unreachable")
}
