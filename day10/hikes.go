package day10

import (
	"aoc24/lib"
	"maps"
)

type TrailPos struct {
	Height  uint8
	Visited bool
}

func WalkTrail(pos lib.WithCoords[TrailPos], grid *lib.Grid[TrailPos]) map[lib.Coords]struct{} {
	nines := make(map[lib.Coords]struct{})
	walkTrail(pos, grid, nines)
	return nines
}

func walkTrail(pos lib.WithCoords[TrailPos], grid *lib.Grid[TrailPos], nines map[lib.Coords]struct{}) {
	pos.Value.Visited = true
	grid.Set(pos.X, pos.Y, pos.Value)

	if pos.Value.Height == 9 {
		nines[pos.Coords] = struct{}{}
	}

	for newpos := range grid.Around(pos.X, pos.Y, lib.Deltas4) {
		if newpos.Value.Visited {
			continue
		}

		if newpos.Value.Height-pos.Value.Height != 1 {
			continue
		}

		walkTrail(newpos, grid, nines)
	}
}

func TrailRating(head, tail lib.WithCoords[TrailPos], grid *lib.Grid[TrailPos]) int {
	return trailRating(head, tail, grid, make(map[lib.Coords]struct{}))
}

func trailRating(pos, target lib.WithCoords[TrailPos], grid *lib.Grid[TrailPos], visited map[lib.Coords]struct{}) (rating int) {
	if pos.Coords == target.Coords {
		return 1
	}

	visited[pos.Coords] = struct{}{}

	for newpos := range grid.Around(pos.X, pos.Y, lib.Deltas4) {
		if _, ok := visited[newpos.Coords]; ok {
			continue
		}

		if newpos.Value.Height-pos.Value.Height != 1 {
			continue
		}

		rating += trailRating(newpos, target, grid, cloneMap(visited))
	}

	return rating
}

func cloneMap[K comparable, V any](data map[K]V) map[K]V {
	newmap := make(map[K]V, len(data))
	maps.Copy(newmap, data)
	return newmap
}
