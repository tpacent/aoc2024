package day20

import (
	"aoc24/lib"
)

type Solver struct {
	track   *lib.Grid[byte]
	visited *lib.Grid[DistInfo]
	vislist []DistCoords
}

type DistInfo struct {
	Distance int
	Visited  bool
}

type DistCoords struct {
	Distance int
	Coords   lib.Coords
}

func NewSolver(track *lib.Grid[byte]) *Solver {
	w, h := track.Width(), track.Height()

	return &Solver{
		track:   track,
		visited: lib.NewGrid(w, h, make([]DistInfo, w*h)),
	}
}

func (s *Solver) Prepare(start, end lib.Coords) {
	s.visited.Set(start.X, start.Y, DistInfo{Visited: true})
	s.race(start, end, 0)

	for _, info := range s.visited.Iter() {
		if info.Value.Visited {
			s.vislist = append(s.vislist, DistCoords{Distance: info.Value.Distance, Coords: info.Coords})
		}
	}
}

var deltas = []lib.Coords{
	{X: 1, Y: 0},
	{X: 0, Y: 1},
	{X: -1, Y: 0},
	{X: 0, Y: -1},
}

func (s *Solver) race(loc, end lib.Coords, dist int) {
	if loc == end {
		return
	}

	for _, next := range deltas {
		next.X += loc.X
		next.Y += loc.Y

		tile, ok := s.track.At(next.X, next.Y)

		if !ok || tile == '#' {
			continue
		}

		info, ok := s.visited.At(next.X, next.Y)
		if !ok {
			continue
		}

		if !info.Visited || info.Distance > dist {
			s.visited.Set(next.X, next.Y, DistInfo{Visited: true, Distance: dist + 1})
			s.race(next, end, dist+1)
		}
	}
}

func (s *Solver) CountCheats(radius, thresold int) (total int) {
	for tile := range s.visited.FindAll(func(wc lib.WithCoords[DistInfo]) bool { return wc.Value.Visited }) {
		for _, target := range s.vislist {
			if tile.Value.Distance < target.Distance {
				continue
			}

			dist := lib.AbsDiff(target.Coords.X, tile.X) + lib.AbsDiff(target.Coords.Y, tile.Y)

			if radius < dist {
				continue
			}

			if save := tile.Value.Distance - target.Distance - dist; save >= thresold {
				total++
			}
		}
	}

	return
}
