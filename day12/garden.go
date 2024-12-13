package day12

import "aoc24/lib"

type Plot struct {
	Type    byte
	Visited bool
}

func FloodFill(pos lib.WithCoords[Plot], grid *lib.Grid[Plot]) (area, perimeter int) {
	area = 1
	perimeter = 4
	pos.Value.Visited = true
	grid.Set(pos.X, pos.Y, pos.Value)
	for nextPos := range grid.Around(pos.X, pos.Y, lib.Deltas4) {
		if nextPos.Value.Type == pos.Value.Type {
			perimeter--
			if !nextPos.Value.Visited {
				a, p := FloodFill(nextPos, grid)
				area += a
				perimeter += p
			}
		}
	}
	return
}
