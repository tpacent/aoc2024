package day25_test

import (
	"aoc24/day25"
	"aoc24/lib"
	"bufio"
	"io"
	"strings"
	"testing"
)

func TestDay25Part1(t *testing.T) {
	file := lib.MustOpenFile("testdata/input.txt")
	t.Cleanup(func() { file.Close() })

	result := day25.NaiveFitPairs(ParseInput(file))
	t.Log(result)
}

const example = `
#####
.####
.####
.####
.#.#.
.#...
.....

#####
##.##
.#.##
...##
...#.
...#.
.....

.....
#....
#....
#...#
#.#.#
#.###
#####

.....
.....
#.#..
###..
###.#
###.#
#####

.....
.....
.....
#....
#.#..
#.#.#
#####
`

func TestParseInput(t *testing.T) {
	r := strings.NewReader(example)
	locks, keys := ParseInput(r)

	t.Log("LOCKS", locks)
	t.Log("KEYS", keys)
}

func ParseInput(r io.Reader) (locks, keys []day25.Pins) {
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		}

		// at first line
		isLock := line[0] == '#'
		pins := day25.Pins{}

		for range 5 {
			_ = scanner.Scan()
			line = scanner.Text()
			for index, c := range line {
				if c == '#' {
					pins[index]++
				}
			}
		}

		_ = scanner.Scan() // skip useless last line

		if isLock {
			locks = append(locks, pins)
		} else {
			keys = append(keys, pins)
		}
	}

	return
}
