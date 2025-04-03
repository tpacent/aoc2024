package day8_test

import (
	"aoc24/day8"
	"aoc24/lib"
	"bufio"
	"io"
	"strings"
	"testing"
)

type Antenna struct {
}

func TestPartOne(t *testing.T) {
	file := lib.MustOpenFile("testdata/input.txt")
	t.Cleanup(func() { _ = file.Close() })

	coordsByType, width, height := fillCoords(file)
	uniqueAntinodes := antinodes(coordsByType, width, height, 1)
	lib.PrintResult(t, 8, 1, len(uniqueAntinodes), 379)
}

func TestPartTwo(t *testing.T) {
	file := lib.MustOpenFile("testdata/input.txt")
	t.Cleanup(func() { _ = file.Close() })

	coordsByType, width, height := fillCoords(file)
	uniqueAntinodes := antinodes(coordsByType, width, height, -1)
	lib.PrintResult(t, 8, 2, len(uniqueAntinodes), 1339)
}

func fillCoords(r io.Reader) (_ map[byte][][2]int, width int, height int) {
	coordsByType := map[byte][][2]int{}
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) == 0 {
			continue
		}
		if width == 0 {
			width = len(line)
		}
		for x, char := range line {
			if char != '.' {
				coordsByType[char] = append(coordsByType[char], [2]int{x, height})
			}
		}
		height++
	}
	return coordsByType, width, height
}

func antinodes(coordsByType map[byte][][2]int, width, height, steps int) map[[2]int]struct{} {
	coordMap := map[[2]int]struct{}{}
	for _, coords := range coordsByType {
		for pair := range day8.Permute(coords) {
			for coord := range day8.PropagateSignal(pair[0], pair[1], steps, width, height) {
				coordMap[coord] = struct{}{}
			}
		}
	}
	return coordMap
}

const example = `
............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`

func TestExample(t *testing.T) {
	coordsByType, width, height := fillCoords(strings.NewReader(example))
	uniqueAntinodes := antinodes(coordsByType, width, height, 1)
	if actual := len(uniqueAntinodes); actual != 14 {
		t.Error("unexpected", actual)
	}
}

const example2 = `
T.........
...T......
.T........
..........
..........
..........
..........
..........
..........
..........`

func TestExample2(t *testing.T) {
	coordsByType, width, height := fillCoords(strings.NewReader(example2))
	uniqueAntinodes := antinodes(coordsByType, width, height, -1)
	if actual := len(uniqueAntinodes); actual != 9 {
		t.Error("unexpected", actual)
	}
}
