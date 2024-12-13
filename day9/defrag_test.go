package day9_test

import (
	"aoc24/day9"
	"aoc24/lib"
	"os"
	"testing"
)

func TestPartOne(t *testing.T) {
	data, err := os.ReadFile("testdata/input.txt")
	if err != nil {
		t.Fatal(err)
	}

	iter := day9.FileIter(ParseInput(string(data)))
	t.Log(day9.Checksum(iter)) // 6288707484810
}

func TestPartTwo(t *testing.T) {
	data, err := os.ReadFile("testdata/input.txt")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(day9.FileIter2(ParseInput(string(data)))) // 6288707484810
}

const example = "2333133121414131402"

func TestExample(t *testing.T) {
	if actual := day9.Checksum(day9.FileIter(ParseInput(example))); actual != 1928 {
		t.Log("unexpected", actual)
	}
}

func TestExample2(t *testing.T) {
	if actual := day9.FileIter2(ParseInput(example)); actual != 2858 {
		t.Log("unexpected", actual)
	}
}

func ParseInput(s string) (frags []day9.Fragment) {
	idCounter := 0

	for index, c := range s {
		isFile := index%2 == 0

		fragment := day9.Fragment{
			IsFile: isFile,
			ID:     idCounter,
			Size:   lib.MustParse(string(c)),
		}

		if !isFile {
			fragment.ID = -1
		}

		frags = append(frags, fragment)

		if isFile {
			idCounter++
		}
	}

	return frags
}
