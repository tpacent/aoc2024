package day16

import (
	"aoc24/lib"
)

type Dir byte

const (
	N Dir = iota
	E
	S
	W
)

type MazeRunner struct {
	maze       *lib.Grid[byte]
	visited    *lib.Grid[int32]
	start, end lib.Coords
	affected   map[lib.Coords]struct{}
	cost       int32
}

func NewMazeRunner(maze *lib.Grid[byte], s, e lib.Coords) *MazeRunner {
	w, h := maze.Width(), maze.Height()

	return &MazeRunner{
		maze:     maze,
		start:    s,
		end:      e,
		visited:  lib.NewGrid(w, h, make([]int32, w*h)),
		affected: make(map[lib.Coords]struct{}),
	}
}

func (mr *MazeRunner) Run() int32 {
	start := StepDir{Coords: mr.start, Dir: E}
	mr.visited.Set(start.X, start.Y, 0)
	mr.run(start, 0, []lib.Coords{mr.start})
	return mr.cost
}

func (mr *MazeRunner) run(step StepDir, cost int32, path []lib.Coords) {
	if step.Coords == mr.end {
		if mr.cost == 0 || cost <= mr.cost {
			mr.fillAffected(path, mr.cost > cost)
			mr.cost = cost
		}
		return
	}

	prevScore := mr.visited.AtUnsafe(step.X, step.Y)
	if prevScore > 0 && cost > prevScore+1000 {
		return
	}

	mr.visited.Set(step.X, step.Y, cost)

	for _, next := range stepDeltas {
		next.X += step.X
		next.Y += step.Y

		if value := mr.maze.AtUnsafe(next.X, next.Y); value == '#' {
			continue
		}

		mr.run(next, cost+mr.stepCost(step.Dir, next.Dir), append(path, next.Coords))
	}
}

func (mr *MazeRunner) fillAffected(path []lib.Coords, overwrite bool) {
	if overwrite {
		clear(mr.affected)
	}

	for _, c := range path {
		mr.affected[c] = struct{}{}
	}
}

func (mr *MazeRunner) AffectedCount() int {
	return len(mr.affected)
}

func (mr *MazeRunner) stepCost(prev, next Dir) int32 {
	switch {
	case prev == next: // forward
		return 1
	case prev+2%4 == next: // opposite direction (2 turns)
		return 2001
	default:
		return 1001 // 1 turn and forward
	}
}

type StepDir struct {
	lib.Coords
	Dir Dir
}

var stepDeltas = []StepDir{
	{Coords: lib.Coords{X: 0, Y: -1}, Dir: N},
	{Coords: lib.Coords{X: 1, Y: 0}, Dir: E},
	{Coords: lib.Coords{X: 0, Y: 1}, Dir: S},
	{Coords: lib.Coords{X: -1, Y: 0}, Dir: W},
}
