package day24_test

import (
	"aoc24/day24"
	"aoc24/lib"
	"bufio"
	"io"
	"strings"
	"testing"
)

func TestDay24Part1(t *testing.T) {
	file := lib.MustOpenFile("testdata/input.txt")
	t.Cleanup(func() { file.Close() })

	regs, gates := ParseInput(file)
	resolver := day24.NewResolver(regs, gates)
	t.Log(day24.SolveZ(resolver)) // 57632654722854
}

func TestDay24Part2(t *testing.T) {
	file := lib.MustOpenFile("testdata/input.txt")
	t.Cleanup(func() { file.Close() })

	regs, gates := ParseInput(file)
	resolver := day24.NewResolver(regs, gates)
	t.Log(strings.Join(day24.FindBrokenGates(resolver), ",")) // ckj,dbp,fdv,kdf,rpp,z15,z23,z39
}

const example = `
x00: 1
x01: 0
x02: 1
x03: 1
x04: 0
y00: 1
y01: 1
y02: 1
y03: 1
y04: 1

ntg XOR fgs -> mjb
y02 OR x01 -> tnw
kwq OR kpj -> z05
x00 OR x03 -> fst
tgd XOR rvg -> z01
vdt OR tnw -> bfw
bfw AND frj -> z10
ffh OR nrd -> bqk
y00 AND y03 -> djm
y03 OR y00 -> psh
bqk OR frj -> z08
tnw OR fst -> frj
gnj AND tgd -> z11
bfw XOR mjb -> z00
x03 OR x00 -> vdt
gnj AND wpb -> z02
x04 AND y00 -> kjc
djm OR pbm -> qhw
nrd AND vdt -> hwm
kjc AND fst -> rvg
y04 OR y02 -> fgs
y01 AND x02 -> pbm
ntg OR kjc -> kwq
psh XOR fgs -> tgd
qhw XOR tgd -> z09
pbm OR djm -> kpj
x03 XOR y03 -> ffh
x00 XOR y04 -> ntg
bfw OR bqk -> z06
nrd XOR fgs -> wpb
frj XOR qhw -> z04
bqk OR frj -> z07
y03 OR x01 -> nrd
hwm AND bqk -> z03
tgd XOR rvg -> z12
tnw OR pbm -> gnj
`

func TestExample(t *testing.T) {
	regs, gates := ParseInput(strings.NewReader(example))
	resolver := day24.NewResolver(regs, gates)
	if actual := day24.SolveZ(resolver); actual != 2024 {
		t.Log("unexpected value")
	}
}

func ParseInput(r io.Reader) (map[string]bool, map[string]day24.GateLogic) {
	registers := map[string]bool{}
	gates := map[string]day24.GateLogic{}
	scanner := bufio.NewScanner(r)
	dataflag := false

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			if dataflag {
				break
			} else {
				continue
			}
		}

		dataflag = true
		reg, val := parseRegister(line)
		registers[reg] = val
	}

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		l, r, op, dst := parseGate(line)
		gates[dst] = day24.GateLogic{L: l, R: r, Op: op}
	}

	return registers, gates
}

func parseRegister(line string) (string, bool) {
	name, sval, _ := strings.Cut(line, ": ")
	return name, sval == "1"
}

func parseGate(line string) (string, string, day24.Op, string) {
	strOp, dst, _ := strings.Cut(line, " -> ")
	parts := strings.Split(strOp, " ")
	return parts[0], parts[2], opMap[parts[1]], dst
}

var opMap = map[string]day24.Op{
	"AND": day24.OpAND,
	"OR":  day24.OpOR,
	"XOR": day24.OpXOR,
}
