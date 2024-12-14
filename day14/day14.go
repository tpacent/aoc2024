package day14

import (
	"aoc24/lib"
)

type Bot struct {
	X, Y   int
	Dx, Dy int
}

type Space struct {
	Bots []*Bot
	W, H int
}

func (s *Space) Run(steps int) {
	for _, bot := range s.Bots {
		x := bot.X + steps*bot.Dx
		y := bot.Y + steps*bot.Dy
		bot.X = lib.Modulo(x, s.W)
		bot.Y = lib.Modulo(y, s.H)
	}
}

func QuadrantCount(s *Space) int {
	quadCounts := [5]int{} // 0 is no quadrant
	quadFunc := quadrant(s.W, s.H)

	for _, bot := range s.Bots {
		quadCounts[quadFunc(bot.X, bot.Y)]++
	}

	return quadCounts[1] * quadCounts[2] * quadCounts[3] * quadCounts[4]
}

func quadrant(w, h int) func(x, y int) int {
	wMid := w / 2 // 1  2
	hMid := h / 2 // 3  4

	return func(x, y int) int {
		switch {
		case x < wMid && y < hMid:
			return 1
		case x > wMid && y < hMid:
			return 2
		case x < wMid && y > hMid:
			return 3
		case x > wMid && y > hMid:
			return 4
		}
		return 0
	}
}
