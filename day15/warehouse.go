package day15

type Coords [2]int

type Warehouse struct {
	Data map[Coords]byte
	Bot  Coords
	W, H int
}

var deltas = map[byte]Coords{
	'<': {-1, 0},
	'^': {0, -1},
	'>': {1, 0},
	'v': {0, 1},
}

func schar(c byte) string {
	switch c {
	case '#', 'O', '@', '[', ']':
		return string(c)
	case 0:
		return "."
	default:
		return "?"
	}
}

func (wh *Warehouse) Move(dir byte) {
	if delta, ok := deltas[dir]; ok {
		if c, ok := Move(wh, wh.Bot, delta, false); ok {
			Move(wh, wh.Bot, delta, true)
			wh.Bot = c
		}
	}
}

func (wh *Warehouse) Print() {
	for y := range wh.H {
		for x := range wh.W {
			coords := [2]int{x, y}

			if coords == wh.Bot {
				print("@")
			} else {
				print(schar(wh.Data[coords]))
			}
		}
		println()
	}
}

func Move(wh *Warehouse, pos Coords, delta Coords, yesmove bool) (Coords, bool) {
	v := wh.Data[pos]

	if v == 0 {
		return [2]int{-1, -1}, true
	}

	if v == '#' {
		return [2]int{-1, -1}, false
	}

	nextPositions := NextPositions(wh.Data, pos, delta)
	for _, nextpos := range nextPositions {
		if _, ok := Move(wh, nextpos, delta, yesmove); !ok {
			return nextpos, false
		}
	}

	if yesmove {
		for _, nextpos := range nextPositions {
			wh.Data[nextpos] = wh.Data[pos]
			delete(wh.Data, pos)
		}
	}

	return nextPositions[0], true
}

func SumGPS(wh *Warehouse, entity byte) (total int) {
	for coords, c := range wh.Data {
		if c != entity {
			continue
		}
		total += coords[0] + 100*coords[1]
	}
	return
}

func NextPositions(data map[Coords]byte, pos, delta [2]int) [][2]int {
	next := [][2]int{{pos[0] + delta[0], pos[1] + delta[1]}}

	if delta[1] != 0 { // move up or down
		switch data[next[0]] {
		case '[':
			next = append(next, [2]int{pos[0] + delta[0] + 1, pos[1] + delta[1]})
		case ']':
			next = append(next, [2]int{pos[0] + delta[0] - 1, pos[1] + delta[1]})
		}
	}

	return next
}
