package day16_test

import (
	"aoc24/day16"
	"aoc24/lib"
	"bufio"
	"io"
	"strings"
	"testing"
)

func TestDay16Part1(t *testing.T) {
	file := lib.MustOpenFile("testdata/input.txt")
	t.Cleanup(func() { _ = file.Close() })
	runner := day16.NewMazeRunner(parseMaze(file))
	t.Log(runner.Run()) // 85396
}

func TestDay16Part2(t *testing.T) {
	file := lib.MustOpenFile("testdata/input.txt")
	t.Cleanup(func() { _ = file.Close() })
	runner := day16.NewMazeRunner(parseMaze(file))
	runner.Run()
	t.Log(runner.AffectedCount()) // 428
}

const example = `
###############
#.......#....E#
#.#.###.#.###.#
#.....#.#...#.#
#.###.#####.#.#
#.#.#.......#.#
#.#.#####.###.#
#...........#.#
###.#.#####.#.#
#...#.....#.#.#
#.#.#.###.#.#.#
#.....#...#.#.#
#.###.#.#.#.#.#
#S..#.....#...#
###############
`

func TestExample(t *testing.T) {
	grid, start, end := parseMaze(strings.NewReader(example))
	runner := day16.NewMazeRunner(grid, start, end)
	if actual := runner.Run(); actual != 7036 {
		t.Error("unexpected value", actual)
	}
	if count := runner.AffectedCount(); count != 45 {
		t.Error("unexpected value", count)
	}
}

const example2 = `
#################
#...#...#...#..E#
#.#.#.#.#.#.#.#.#
#.#.#.#...#...#.#
#.#.#.#.###.#.#.#
#...#.#.#.....#.#
#.#.#.#.#.#####.#
#.#...#.#.#.....#
#.#.#####.#.###.#
#.#.#.......#...#
#.#.###.#####.###
#.#.#...#.....#.#
#.#.#.#####.###.#
#.#.#.........#.#
#.#.#.#########.#
#S#.............#
#################
`

func TestExample2(t *testing.T) {
	grid, start, end := parseMaze(strings.NewReader(example2))
	runner := day16.NewMazeRunner(grid, start, end)
	if actual := runner.Run(); actual != 11048 {
		t.Error("unexpected value", actual)
	}
	if count := runner.AffectedCount(); count != 64 {
		t.Error("unexpected value", count)
	}
}

func parseMaze(input io.Reader) (*lib.Grid[byte], lib.Coords, lib.Coords) {
	var width, height int
	var start, end lib.Coords

	scanner := bufio.NewScanner(input)

	data := make([]byte, 0)

	for scanner.Scan() {
		line := scanner.Bytes()

		if len(line) == 0 {
			continue // ignore empty start/end
		}

		width = len(line)

		for index, c := range line {
			if c == '#' {
				data = append(data, '#')
			} else {
				data = append(data, '.')
			}

			switch c {
			case 'S':
				start = lib.Coords{X: index, Y: height}
			case 'E':
				end = lib.Coords{X: index, Y: height}
			}
		}

		height++
	}

	return lib.NewGrid(width, height, data), start, end
}
