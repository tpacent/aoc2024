package day14

import (
	"aoc24/lib"
	"fmt"
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

func QuadrantCount(s *Space) [5]int {
	quadCounts := [5]int{} // 0 is no quadrant
	quadFunc := quadrant(s.W, s.H)

	for _, bot := range s.Bots {
		quadCounts[quadFunc(bot.X, bot.Y)]++
	}

	return quadCounts
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

func SignalDetect(w, h, runlen int) func(bots []*Bot) bool {
	botmap := make(map[[2]int]int, 1024)

	return func(bots []*Bot) bool {
		clear(botmap)

		for _, bot := range bots {
			botmap[[2]int{bot.X, bot.Y}]++
		}

		for y := 0; y < h; y++ {
			seqCurr := 0
			for x := 0; x < w; x++ {
				if botmap[[2]int{x, y}] == 0 {
					seqCurr = 0
					continue
				}
				seqCurr++
				if seqCurr >= runlen {
					return true
				}
			}
		}

		return false
	}
}

func PrintSpace(space *Space) {
	botmap := map[[2]int]int{}
	for _, bot := range space.Bots {
		botmap[[2]int{bot.X, bot.Y}]++
	}

	for y := 0; y < space.H; y++ {
		for x := 0; x < space.W; x++ {
			if n := botmap[[2]int{x, y}]; n == 0 {
				fmt.Print(".")
			} else {
				fmt.Print(n)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
