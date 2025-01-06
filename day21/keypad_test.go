package day21_test

import (
	"aoc24/day21"
	"aoc24/lib"
	"bufio"
	"io"
	"strings"
	"testing"
)

func TestDay21Part1(t *testing.T) {
	file := lib.MustOpenFile("testdata/input.txt")
	t.Cleanup(func() { _ = file.Close() })
	total := 0
	for _, code := range parseInput(file) {
		keypad := day21.NewStackPad(2)
		total += day21.Complexity(keypad, code)
	}
	t.Log(total) // 270084
}

func parseInput(r io.Reader) (out [][]byte) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) == 0 {
			continue
		}
		out = append(out, line)
	}
	return
}

func TestExamples(t *testing.T) {
	tcases := []struct {
		Input    []byte
		Expected string
	}{
		{Input: []byte{'0', '2', '9', 'A'}, Expected: "<vA<AA>>^AvAA<^A>A<v<A>>^AvA^A<vA>^A<v<A>^A>AAvA^A<v<A>A>^AAAvA<^A>A"},
		{Input: []byte{'9', '8', '0', 'A'}, Expected: "<v<A>>^AAAvA^A<vA<AA>>^AvAA<^A>A<v<A>A>^AAAvA<^A>A<vA>^A<A>A"},
		{Input: []byte{'1', '7', '9', 'A'}, Expected: "<v<A>>^A<vA<A>>^AAvAA<^A>A<v<A>>^AAvA^A<vA>^AA<A>A<v<A>A>^AAAvA<^A>A"},
		{Input: []byte{'4', '5', '6', 'A'}, Expected: "<v<A>>^AA<vA<A>>^AAvAA<^A>A<vA>^A<A>A<vA>^A<A>A<v<A>A>^AAvA<^A>A"},
		{Input: []byte{'3', '7', '9', 'A'}, Expected: "<v<A>>^AvA^A<vA<AA>>^AAvA<^A>AAvA^A<vA>^AA<A>A<v<A>A>^AAAvA<^A>A"},
	}

	complexity := 0

	for _, tt := range tcases {
		t.Run(string(tt.Input), func(t *testing.T) {
			sb := strings.Builder{}

			keypad := day21.NewStackPad(2)
			for _, btn := range tt.Input {
				for _, chunk := range keypad.Dial(btn) {
					sb.WriteString(string(chunk))
					break
				}
			}

			if actual := sb.String(); len(actual) != len(tt.Expected) {
				t.Errorf(`unexpected value "%s"; expected "%s"`, actual, tt.Expected)
			}

			complexity += day21.Complexity(keypad, tt.Input)
		})
	}

	if complexity != 126384 {
		t.Error("unexpected total")
	}
}
