package day4_test

import (
	"aoc24/day4"
	"aoc24/lib"
	"bufio"
	"bytes"
	"errors"
	"io"
	"strings"
	"testing"
)

func TestPartOne(t *testing.T) {
	file := lib.MustOpenFile("testdata/input.txt")
	t.Cleanup(func() { _ = file.Close() })
	grid := lib.NewGrid(readGrid(file))
	t.Log(day4.SearchDir(grid, []byte("XMAS"))) // 2639
}

func TestPartTwo(t *testing.T) {
	file := lib.MustOpenFile("testdata/input.txt")
	t.Cleanup(func() { _ = file.Close() })
	grid := lib.NewGrid(readGrid(file))
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
	grid := lib.NewGrid(readGrid(strings.NewReader(example)))
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
	grid := lib.NewGrid(readGrid(strings.NewReader(example2)))
	if actual := day4.SearchXMas(grid); actual != 9 {
		t.Error("unexpected value", actual)
	}
}

func readGrid(src io.Reader) (w, h int, data []byte) {
	r := bufio.NewReader(src)

	for {
		line, err := r.ReadBytes('\n')
		chunk := bytes.TrimSuffix(line, []byte{'\n'})

		if len(chunk) > 0 {
			data = append(data, chunk...)
			h++
		}

		if errors.Is(err, io.EOF) {
			break
		}
	}

	w = len(data) / h
	return
}
