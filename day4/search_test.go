package day4_test

import (
	"aoc24/day4"
	"aoc24/lib"
	"strings"
	"testing"
)

func byteGrid(b byte) byte { return b }

func TestPartOne(t *testing.T) {
	file := lib.MustOpenFile("testdata/input.txt")
	t.Cleanup(func() { _ = file.Close() })
	grid := lib.NewGrid(lib.ReadGrid(file, byteGrid))
	t.Log(day4.SearchDir(grid, []byte("XMAS"))) // 2639
}

func TestPartTwo(t *testing.T) {
	file := lib.MustOpenFile("testdata/input.txt")
	t.Cleanup(func() { _ = file.Close() })
	grid := lib.NewGrid(lib.ReadGrid(file, byteGrid))
	t.Log(day4.SearchXMas(grid)) // 2005
}

const example = `
MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

func TestExample(t *testing.T) {
	grid := lib.NewGrid(lib.ReadGrid(strings.NewReader(example), byteGrid))
	if actual := day4.SearchDir(grid, []byte("XMAS")); actual != 18 {
		t.Error("unexpected value", actual)
	}
}

const example2 = `
.M.S......
..A..MSMS.
.M.S.MAA..
..A.ASMSM.
.M.S.M....
..........
S.S.S.S.S.
.A.A.A.A..
M.M.M.M.M.
..........`

func TestExample2(t *testing.T) {
	grid := lib.NewGrid(lib.ReadGrid(strings.NewReader(example2), byteGrid))
	if actual := day4.SearchXMas(grid); actual != 9 {
		t.Error("unexpected value", actual)
	}
}
