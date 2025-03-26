package day23_test

import (
	"aoc24/day23"
	"aoc24/lib"
	"bufio"
	"io"
	"strings"
	"testing"
)

func TestDay23Part1(t *testing.T) {
	file := lib.MustOpenFile("testdata/input.txt")
	t.Cleanup(func() { file.Close() })

	graph := ParseInput(file)
	predicate := func(s string) bool { return s[0] == 't' }
	t.Log(day23.CountLoops(graph, predicate)) // 1323
}

func TestDay23Part2(t *testing.T) {
	file := lib.MustOpenFile("testdata/input.txt")
	t.Cleanup(func() { file.Close() })

	t.Log(day23.GetPassword(ParseInput(file))) // er,fh,fi,ir,kk,lo,lp,qi,ti,vb,xf,ys,yu
}

func ParseInput(r io.Reader) *day23.Graph {
	graph := day23.NewGraph()

	for s := bufio.NewScanner(r); s.Scan(); {
		line := strings.TrimSpace(s.Text())
		a, b, ok := strings.Cut(line, "-")
		if !ok {
			continue
		}
		graph.Add(a, b)
	}

	return graph
}

const example = `
kh-tc
qp-kh
de-cg
ka-co
yn-aq
qp-ub
cg-tb
vc-aq
tb-ka
wh-tc
yn-cg
kh-ub
ta-co
de-co
tc-td
tb-wq
wh-td
ta-ka
td-qp
aq-cg
wq-ub
ub-vc
de-ta
wq-aq
wq-vc
wh-yn
ka-de
kh-ta
co-tc
wh-qp
tb-vc
td-yn
`

func TestExample(t *testing.T) {
	graph := ParseInput(strings.NewReader(example))
	predicate := func(s string) bool { return s[0] == 't' }
	if actual := day23.CountLoops(graph, predicate); actual != 7 {
		t.Errorf("unexpected value: %d", actual)
	}
}

func TestExample2(t *testing.T) {
	graph := ParseInput(strings.NewReader(example))
	if password := day23.GetPassword(graph); password != "co,de,ka,ta" {
		t.Error("unexpected value")
	}
}
