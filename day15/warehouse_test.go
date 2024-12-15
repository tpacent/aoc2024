package day15_test

import (
	"aoc24/day15"
	"bufio"
	"iter"
	"os"
	"strings"
	"testing"
)

func TestPartOne(t *testing.T) {
	data, _ := os.ReadFile("testdata/input.txt")
	wh, it := parseInput(string(data), false)
	for dir := range it {
		wh.Move(dir)
	}

	t.Log(day15.SumGPS(wh, 'O')) // 1360570
}

func TestPartTwo(t *testing.T) {
	data, _ := os.ReadFile("testdata/input.txt")
	wh, it := parseInput(string(data), true)
	for dir := range it {
		wh.Move(dir)
	}

	t.Log(day15.SumGPS(wh, '[')) // 1381446
}

const example = `
##########
#..O..O.O#
#......O.#
#.OO..O.O#
#..O@..O.#
#O#..O...#
#O..O..O.#
#.OO.O.OO#
#....O...#
##########

<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^
vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v
><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<
<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^
^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><
^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^
>^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^
<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>
^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>
v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^
`

func TestExample(t *testing.T) {
	wh, it := parseInput(example, false)
	for dir := range it {
		wh.Move(dir)
	}
	wh.Print()
}

const example2 = `
#######
#...#.#
#.....#
#..OO@#
#..O..#
#.....#
#######

<vv<<^^<<^^
`

func TestExample2(t *testing.T) {
	wh, it := parseInput(example2, true)
	for dir := range it {
		wh.Move(dir)
	}
	wh.Print()
}

func parseInput(input string, expand bool) (*day15.Warehouse, iter.Seq[byte]) {
	grid, moves, _ := strings.Cut(input, "\n\n")

	data := make(map[day15.Coords]byte, len(grid))
	w, h := 0, 0
	bot := day15.Coords{}

	scanner := bufio.NewScanner(strings.NewReader(grid))
	for scanner.Scan() {
		line := scanner.Bytes()

		if len(line) == 0 {
			continue
		}

		h++
		w = len(line)

		for index, c := range line {
			if expand {
				coords := [2]int{index * 2, h - 1}
				switch c {
				case '@':
					data[coords] = c
					bot = coords
				case 'O':
					data[coords] = '['
					data[[2]int{coords[0] + 1, coords[1]}] = ']'
				case '#':
					data[coords] = c
					data[[2]int{coords[0] + 1, coords[1]}] = c
				}
			} else {
				coords := [2]int{index, h - 1}
				switch c {
				case '@':
					data[coords] = c
					bot = coords
				case '#', 'O', '[', ']':
					data[coords] = c
				}
			}
		}
	}

	if expand {
		w *= 2
	}

	moveIter := func(yield func(byte) bool) {
		r := strings.NewReader(moves)

		for {
			b, err := r.ReadByte()
			if err != nil {
				return
			}
			switch b {
			case '<', '>', 'v', '^':
				if ok := yield(b); !ok {
					return
				}
			}
		}
	}

	return &day15.Warehouse{
		Data: data,
		Bot:  bot,
		W:    w,
		H:    h,
	}, moveIter
}
