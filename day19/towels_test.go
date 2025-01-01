package day19_test

import (
	"aoc24/day19"
	"aoc24/lib"
	"bufio"
	"io"
	"strings"
	"testing"
)

func TestDay19(t *testing.T) {
	file := lib.MustOpenFile("testdata/input.txt")
	t.Cleanup(func() { _ = file.Close() })
	patterns, designs := parseInput(file)
	patternSet := lib.MakeSet(patterns)
	availableDesigns, combinations := 0, 0
	for _, design := range designs {
		if value := day19.CountArrangements(design, patternSet); value > 0 {
			availableDesigns++
			combinations += value
		}
	}

	t.Log("designs", availableDesigns) // 308
	t.Log("ways total", combinations)  // 662726441391898
}

const example = `
r, wr, b, g, bwu, rb, gb, br

brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb
`

func TestExample(t *testing.T) {
	patterns, designs := parseInput(strings.NewReader(example))

	passed, combs := 0, 0
	for _, design := range designs {
		if value := day19.CountArrangements(design, lib.MakeSet(patterns)); value > 0 {
			passed++
			combs += value
		}
	}
	if passed != 6 {
		t.Error("unexpected passing designs")
	}
	if combs != 16 {
		t.Error("unexpected combinations")
	}
}

func parseInput(r io.Reader) (patterns []string, designs []string) {
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		if line := scanner.Text(); len(line) > 0 {
			patterns = strings.Split(line, ", ")
			break
		}
	}

	for scanner.Scan() {
		if line := scanner.Text(); len(line) > 0 {
			designs = append(designs, line)
		}
	}

	return
}
