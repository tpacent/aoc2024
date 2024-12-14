package day14_test

import (
	"aoc24/day14"
	"aoc24/lib"
	"bufio"
	"io"
	"strings"
	"testing"
)

func TestPartOne(t *testing.T) {
	file := lib.MustOpenFile("testdata/input.txt")
	t.Cleanup(func() { _ = file.Close() })
	bots := parseInput(file)
	space := &day14.Space{Bots: bots, W: 101, H: 103}
	space.Run(100)
	v := day14.QuadrantCount(space)
	t.Log(v[1] * v[2] * v[3] * v[4]) // 225943500
}

func TestPartTwo(t *testing.T) {
	file := lib.MustOpenFile("testdata/input.txt")
	t.Cleanup(func() { _ = file.Close() })
	bots := parseInput(file)
	space := &day14.Space{Bots: bots, W: 101, H: 103}
	detector := day14.SignalDetect(space.W, space.H, 16)
	n := 0
	for range 10403 { // period is 10403 (103 * 101)
		n++
		space.Run(1)
		if detector(space.Bots) {
			break
		}
	}

	day14.PrintSpace(space)
	t.Log(n) // 6377
}

// 1111111111111111111111111111111
// 1.............................1
// 1.............................1
// 1.............................1
// 1.............................1
// 1..............1..............1
// 1.............111.............1
// 1............11111............1
// 1...........1111111...........1
// 1..........111111111..........1
// 1............11111............1
// 1...........1111111...........1
// 1..........111111111..........1
// 1.........11111111111.........1
// 1........1111111111111........1
// 1..........111111111..........1
// 1.........11111111111.........1
// 1........1111111111111........1
// 1.......111111111111111.......1
// 1......11111111111111111......1
// 1........1111111111111........1
// 1.......111111111111111.......1
// 1......11111111111111111......1
// 1.....1111111111111111111.....1
// 1....111111111111111111111....1
// 1.............111.............1
// 1.............111.............1
// 1.............111.............1
// 1.............................1
// 1.............................1
// 1.............................1
// 1.............................1
// 1111111111111111111111111111111

const example = `
p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3`

func TestExample(t *testing.T) {
	bots := parseInput(strings.NewReader(example))
	space := &day14.Space{
		Bots: bots,
		W:    11,
		H:    7,
	}
	space.Run(100)
	v := day14.QuadrantCount(space)
	if actual := v[1] * v[2] * v[3] * v[4]; actual != 12 {
		t.Error("unexpected", actual)
	}
}

func parseInput(r io.Reader) (bots []*day14.Bot) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		px, py, dx, dy := parseLine(line)
		bot := day14.Bot{X: px, Y: py, Dx: dx, Dy: dy}
		bots = append(bots, &bot)
	}
	return
}

func parseLine(line string) (int, int, int, int) {
	fields := strings.Fields(line)
	_, ps, _ := strings.Cut(fields[0], "=")
	pn := strings.Split(ps, ",")
	_, vs, _ := strings.Cut(fields[1], "=")
	vn := strings.Split(vs, ",")
	return lib.MustParse(pn[0]),
		lib.MustParse(pn[1]),
		lib.MustParse(vn[0]),
		lib.MustParse(vn[1])
}
