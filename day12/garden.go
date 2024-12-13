package day12

import "aoc24/lib"

type Plot struct {
	Type    byte
	Visited bool
}

const (
	Hrz = 1
	Vrt = 2
)

const (
	FromLeft = 1 + iota
	FromRight
	FromAbove
	FromBelow
)

type Wall struct {
	Dir uint8
	lib.Coords
}

func FloodFill(pos lib.WithCoords[Plot], grid *lib.Grid[Plot]) (int, map[Wall]uint8) {
	area := 1

	perimeter := map[Wall]uint8{
		{Dir: Hrz, Coords: pos.Coords}:                         FromBelow,
		{Dir: Vrt, Coords: pos.Coords}:                         FromRight,
		{Dir: Hrz, Coords: lib.Coords{X: pos.X, Y: pos.Y + 1}}: FromAbove,
		{Dir: Vrt, Coords: lib.Coords{X: pos.X + 1, Y: pos.Y}}: FromLeft,
	}

	pos.Value.Visited = true
	grid.Set(pos.X, pos.Y, pos.Value)
	for nextPos := range grid.Around(pos.X, pos.Y, lib.Deltas4) {
		if nextPos.Value.Type != pos.Value.Type {
			continue
		}

		delete(perimeter, getWall(pos.Coords, nextPos.Coords))

		if !nextPos.Value.Visited {
			a, perms := FloodFill(nextPos, grid)
			area += a
			for p, from := range perms {
				perimeter[p] = from
			}
		}
	}
	return area, perimeter
}

func getWall(curr, next lib.Coords) Wall {
	switch {
	case next.X > curr.X:
		return Wall{Dir: Vrt, Coords: next}
	case next.X < curr.X:
		return Wall{Dir: Vrt, Coords: curr}
	case next.Y > curr.Y:
		return Wall{Dir: Hrz, Coords: next}
	case next.Y < curr.Y:
		return Wall{Dir: Hrz, Coords: curr}
	}

	panic("unreachable")
}

func CalcTotalPerimeter(grid *lib.Grid[Plot]) (total int) {
	for _, cell := range grid.Iter() {
		if cell.Value.Visited {
			continue
		}
		area, perimeter := FloodFill(cell, grid)
		total += area * len(perimeter)
	}
	return
}

func CalcTotalWalls(grid *lib.Grid[Plot]) (total int) {
	for _, cell := range grid.Iter() {
		if cell.Value.Visited {
			continue
		}
		area, perimeter := FloodFill(cell, grid)
		total += area * calcNumWalls(perimeter)
	}
	return
}

var adjacentDeltas = map[uint8][][2]int{
	Hrz: {{-1, 0}, {1, 0}},
	Vrt: {{0, -1}, {0, 1}},
}

func calcNumWalls(perim map[Wall]uint8) (count int) {
	for len(perim) > 0 {
		wall, firstFrom, _ := getOne(perim)
		delete(perim, wall)
		count++

		deltas := adjacentDeltas[wall.Dir]

		for _, delta := range deltas {
			coords := wall.Coords

			for {
				coords.X += delta[0]
				coords.Y += delta[1]
				wall := Wall{Dir: wall.Dir, Coords: coords}

				if from, ok := perim[wall]; ok {
					if from == firstFrom {
						delete(perim, wall)
						continue
					}
				}

				break
			}
		}
	}

	return count
}

func getOne[K comparable, V any](items map[K]V) (K, V, bool) {
	if len(items) == 0 {
		var k K
		var v V
		return k, v, false
	}

	for k, v := range items {
		return k, v, true
	}

	panic("unreachable")
}
