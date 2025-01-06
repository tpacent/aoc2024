package day21

import (
	"aoc24/lib"
	"bytes"
	"cmp"
)

func NewStackPad(dirpads int) *Keypad {
	keypad := MakeKeypad()
	for range dirpads {
		keypad = MakeDirpad(keypad)
	}
	return keypad
}

func Complexity(keypad *Keypad, code []byte) int {
	num := lib.MustParse(string(code[:3]))
	seq := make([]byte, 0)
	for _, btn := range code {
		var shortestChunk []byte
		for _, chunk := range keypad.Dial(btn) {
			if shortestChunk == nil || len(chunk) < len(shortestChunk) {
				shortestChunk = chunk
			}
		}
		seq = append(seq, shortestChunk...)
	}
	return num * len(seq)
}

type Keypad struct {
	state     lib.Coords
	buttons   map[byte]lib.Coords
	reachable map[lib.Coords]struct{}
	childpad  *Keypad
}

func (k *Keypad) Loc() lib.Coords {
	return k.state
}

func (k *Keypad) SetLoc(loc lib.Coords) {
	k.state = loc
}

type MoveSpec struct {
	lib.Coords
	Step byte
}

func (k *Keypad) Move(loc lib.Coords, to byte) (lib.Coords, [][]byte) {
	dst, ok := k.buttons[to]
	if !ok {
		panic("unreachable")
	}

	out := make([][]byte, 0, 2)

	if _, moves := k.moveXthenY(loc, dst); k.validMoves(moves) {
		out = append(out, extractSteps(moves))
	}

	if _, moves := k.moveYthenX(loc, dst); k.validMoves(moves) {
		steps := extractSteps(moves)
		if len(out) == 0 || !bytes.Equal(steps, out[0]) {
			out = append(out, steps)
		}
	}

	k.SetLoc(dst)

	return dst, out
}

func (k *Keypad) validMoves(steps []MoveSpec) bool {
	for _, step := range steps {
		if _, ok := k.reachable[step.Coords]; !ok {
			return false
		}
	}

	return true
}

func (k *Keypad) moveXthenY(loc, dst lib.Coords) (lib.Coords, []MoveSpec) {
	intermediate, movesX := k.moveX(loc, dst)
	_, movesY := k.moveY(intermediate, dst)
	return dst, append(movesX, movesY...)
}

func (k *Keypad) moveYthenX(loc, dst lib.Coords) (lib.Coords, []MoveSpec) {
	intermediate, movesY := k.moveY(loc, dst)
	_, movesX := k.moveX(intermediate, dst)
	return dst, append(movesY, movesX...)
}

func (k *Keypad) moveX(loc, dst lib.Coords) (lib.Coords, []MoveSpec) {
	var step byte
	moves := make([]MoveSpec, 0, lib.AbsDiff(loc.X, dst.X))
	dx := cmp.Compare(dst.X, loc.X)
	switch dx {
	case 1:
		step = '>'
	case -1:
		step = '<'
	}
	for {
		if loc.X == dst.X {
			break
		}
		loc.X += dx
		moves = append(moves, MoveSpec{Coords: loc, Step: step})
	}
	return loc, moves
}

func (k *Keypad) moveY(loc, dst lib.Coords) (lib.Coords, []MoveSpec) {
	var step byte
	moves := make([]MoveSpec, 0, lib.AbsDiff(loc.Y, dst.Y))
	dy := cmp.Compare(dst.Y, loc.Y)
	switch dy {
	case 1:
		step = 'v'
	case -1:
		step = '^'
	}
	for {
		if loc.Y == dst.Y {
			break
		}
		loc.Y += dy
		moves = append(moves, MoveSpec{Coords: loc, Step: step})
	}
	return loc, moves
}

func (k *Keypad) Dial(btn byte) (out [][]byte) {
	if k.childpad == nil { // last one in chain
		_, paths := k.Move(k.Loc(), btn)
		for _, path := range paths {
			moves := append(path, 'A')
			out = append(out, moves)
		}
		return
	}

	paths := k.childpad.Dial(btn)
	for _, path := range paths {
		var shortestPath []byte
		for _, btn := range path {
			var shortestBranch []byte
			_, moveopts := k.Move(k.Loc(), btn)
			for _, branch := range moveopts {
				if shortestBranch == nil || len(branch) < len(shortestBranch) {
					shortestBranch = append(branch, 'A')
				}
			}
			shortestPath = append(shortestPath, shortestBranch...)
		}
		out = append(out, shortestPath)
	}

	return out
}

func NewKeypad(b map[byte]lib.Coords, init byte, child *Keypad) *Keypad {
	initloc := b[init]

	reachable := make(map[lib.Coords]struct{})
	for _, coord := range b {
		reachable[coord] = struct{}{}
	}

	return &Keypad{
		state:     initloc,
		buttons:   b,
		childpad:  child,
		reachable: reachable,
	}
}

func MakeKeypad() *Keypad {
	return NewKeypad(map[byte]lib.Coords{
		'7': {X: 0, Y: 0}, '8': {X: 1, Y: 0}, '9': {X: 2, Y: 0},
		'4': {X: 0, Y: 1}, '5': {X: 1, Y: 1}, '6': {X: 2, Y: 1},
		'1': {X: 0, Y: 2}, '2': {X: 1, Y: 2}, '3': {X: 2, Y: 2},
		'0': {X: 1, Y: 3}, 'A': {X: 2, Y: 3},
	}, 'A', nil)
}

func MakeDirpad(keypad *Keypad) *Keypad {
	return NewKeypad(map[byte]lib.Coords{
		'^': {X: 1, Y: 0}, 'A': {X: 2, Y: 0},
		'<': {X: 0, Y: 1}, 'v': {X: 1, Y: 1}, '>': {X: 2, Y: 1},
	}, 'A', keypad)
}

func extractSteps(moves []MoveSpec) []byte {
	steps := make([]byte, 0, len(moves))
	for _, move := range moves {
		steps = append(steps, move.Step)
	}
	return steps
}
