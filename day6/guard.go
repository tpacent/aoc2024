package day6

import "aoc24/lib"

type Dir uint8

const (
	DirUnknown = iota
	DirU
	DirD
	DirL
	DirR
)

var dirs = map[Dir][2]int{
	DirU: {0, -1},
	DirD: {0, 1},
	DirL: {-1, 0},
	DirR: {1, 0},
}

type Cell struct {
	Empty   bool
	Visited bool
}

type State struct {
	X, Y int
	Dir  Dir
}

func NewWalker(grid *lib.Grid[Cell], state State) *Walker {
	return &Walker{
		grid:  grid,
		state: state,
	}
}

type Walker struct {
	grid  *lib.Grid[Cell]
	state State
}

func (w *Walker) Walk() int {
	x, y := w.state.X, w.state.Y
	visited := map[[2]int]struct{}{
		{x, y}: {},
	}

	for {
		dirDeltas := dirs[w.state.Dir]
		nx, ny := x+dirDeltas[0], y+dirDeltas[1]
		nextCell, ok := w.grid.At(nx, ny)

		if !ok {
			break // outside grid
		}

		if !nextCell.Empty {
			// turn right
			switch w.state.Dir {
			case DirU:
				w.state.Dir = DirR
			case DirR:
				w.state.Dir = DirD
			case DirD:
				w.state.Dir = DirL
			case DirL:
				w.state.Dir = DirU
			}
			continue
		}

		x = nx
		y = ny
		visited[[2]int{x, y}] = struct{}{}
	}

	return len(visited)
}
