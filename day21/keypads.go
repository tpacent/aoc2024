package day21

import (
	"aoc24/lib"
	"cmp"
	"slices"
	"strings"
)

type Coords [2]int8

const (
	MoveNone byte = 0
	MoveUp   byte = '^'
	MoveRt   byte = '>'
	MoveDn   byte = 'v'
	MoveLt   byte = '<'
)

var defaultPriorities = []byte{MoveLt, MoveDn, MoveUp, MoveRt}

type MoveAxis byte

const (
	AxisNone MoveAxis = iota
	AxisH
	AxisV
)

type MDir [2]int8

var moveDeltas = map[byte]MDir{
	MoveUp: {0, -1},
	MoveDn: {0, 1},
	MoveLt: {-1, 0},
	MoveRt: {1, 0},
}

var moveAxes = map[byte]MoveAxis{
	MoveUp: AxisV,
	MoveDn: AxisV,
	MoveLt: AxisH,
	MoveRt: AxisH,
}

func NewKeypad(ks map[byte]Coords, mp []byte) *KeyPad {
	coords := make(map[Coords]byte)
	for key, coord := range ks {
		coords[coord] = key
	}

	return &KeyPad{
		keys:         ks,
		coords:       coords,
		movePriority: mp,
	}
}

type KeyPad struct {
	keys         map[byte]Coords
	coords       map[Coords]byte
	movePriority []byte
}

func (kp *KeyPad) MoveBtn(loc, tgt byte) []byte {
	return kp.Move(kp.keys[loc], kp.keys[tgt])
}

func (kp *KeyPad) Move(loc, tgt Coords) []byte {
	if loc == tgt {
		return nil
	}

	movePlan := map[byte]int8{}

	switch cmp.Compare(loc[0], tgt[0]) {
	case -1:
		movePlan[MoveRt] = lib.AbsDiff(loc[0], tgt[0])
	case 1:
		movePlan[MoveLt] = lib.AbsDiff(loc[0], tgt[0])
	}

	switch cmp.Compare(loc[1], tgt[1]) {
	case -1:
		movePlan[MoveDn] = lib.AbsDiff(loc[1], tgt[1])
	case 1:
		movePlan[MoveUp] = lib.AbsDiff(loc[1], tgt[1])
	}

	axesOrder := []MoveAxis{}
	hMoves := []byte{}
	vMoves := []byte{}
	moveBuf := make([]byte, 0)

	// consider moves in order of priority
	for _, move := range kp.movePriority {
		count := movePlan[move]

		if count == 0 {
			continue
		}

		moveBuf = moveBuf[:0]
		axis := moveAxes[move]
		axesOrder = append(axesOrder, axis)

		for range count {
			moveBuf = append(moveBuf, move)
		}

		switch axis {
		case AxisH:
			hMoves = append(hMoves, moveBuf...)
		case AxisV:
			vMoves = append(vMoves, moveBuf...)
		}
	}

	moves := make([]byte, 0, len(hMoves)+len(vMoves))

	for _, axis := range axesOrder {
		switch axis {
		case AxisH:
			moves = append(moves, hMoves...)
		case AxisV:
			moves = append(moves, vMoves...)
		}
	}

	if !kp.checkMoves(loc, moves) {
		slices.Reverse(moves)
	}

	return moves
}

func (kp *KeyPad) checkMoves(loc Coords, moves []byte) bool {
	for _, move := range moves {
		delta := moveDeltas[move]
		loc[0] += delta[0]
		loc[1] += delta[1]
		if _, ok := kp.coords[loc]; !ok {
			return false
		}
	}
	return true
}

type ExpBag struct {
	Pad   *KeyPad
	Lvl   int
	Cache []map[string]int
}

func ExpandMoves(pad *KeyPad, moves []byte, depth int) int {
	cache := make([]map[string]int, depth+1)
	for i := range cache {
		cache[i] = map[string]int{}
	}

	return expandMoves(ExpBag{
		Pad:   pad,
		Lvl:   depth,
		Cache: cache,
	}, moves)
}

func expandMoves(info ExpBag, moves []byte) (total int) {
	if value, ok := info.Cache[info.Lvl][string(moves)]; ok {
		return value
	}

	if info.Lvl == 0 {
		total = len(moves)
		info.Cache[0][string(moves)] = total
		return
	}

	loc := byte('A') // initially aim at the keypad's A key

	for _, key := range moves {
		total += expandMoves(ExpBag{
			Pad:   info.Pad,
			Lvl:   info.Lvl - 1,
			Cache: info.Cache,
		}, append(info.Pad.MoveBtn(loc, key), 'A'))
		loc = key
	}

	info.Cache[info.Lvl][string(moves)] = total
	return
}

func CodeComplexity(code []byte, dirLevels int) int {
	keypad := NewKeypad(NumPadLayout, defaultPriorities)
	dirpad := NewKeypad(DirPadLayout, defaultPriorities)

	keypadMoves := []byte{}
	loc := byte('A')
	for _, key := range code {
		keypadMoves = append(keypadMoves, keypad.MoveBtn(loc, key)...)
		keypadMoves = append(keypadMoves, 'A')
		loc = key
	}

	codenum := lib.MustParse(strings.TrimRight(string(code), "A"))
	return codenum * ExpandMoves(dirpad, keypadMoves, dirLevels)
}
