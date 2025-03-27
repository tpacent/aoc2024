package day24

import (
	"aoc24/lib"
	"fmt"
	"slices"
)

type Op uint8

const (
	OpUnspecified Op = iota
	OpAND
	OpOR
	OpXOR
)

type GateLogic struct {
	L  string
	R  string
	Op Op
}

func exec(l, r bool, op Op) bool {
	switch op {
	case OpAND:
		return l && r
	case OpOR:
		return l || r
	case OpXOR:
		return l != r
	}

	panic("unreachable")
}

func NewResolver(regs map[string]bool, gates map[string]GateLogic) *Resolver {
	return &Resolver{
		registers: regs,
		gates:     gates,
	}
}

type Resolver struct {
	registers map[string]bool
	gates     map[string]GateLogic
}

func (r *Resolver) Resolve(key string) bool {
	if value, ok := r.registers[key]; ok {
		return value
	}

	gate := r.gates[key]
	value := exec(
		r.Resolve(gate.L),
		r.Resolve(gate.R),
		gate.Op,
	)
	r.registers[key] = value
	return value
}

func (r *Resolver) Match(f func(GateLogic) bool) string {
	for dst, gl := range r.gates {
		if f(gl) {
			return dst
		}
	}

	return ""
}

func SolveZ(resolver *Resolver) (n int) {
	for k := range resolver.gates {
		if k[0] != 'z' {
			continue
		}

		if resolver.Resolve(k) {
			offset := lib.MustParse(k[1:])
			n |= 1 << offset
		}
	}

	return
}

type Adder struct {
	Index     int
	X, Y      string
	Z         string
	Cin, Cout string
	XYxor     string
	XYand     string
	Cand      string
}

func MakeGate(r *Resolver, index int) *Adder {
	x := fmt.Sprintf("x%02d", index)
	y := fmt.Sprintf("y%02d", index)

	gate := Adder{Index: index, X: x, Y: y}

	// huge ugly code below

	if dst := r.Match(func(gl GateLogic) bool {
		return gl.Op == OpXOR && ((x == gl.L && y == gl.R) || (y == gl.L && x == gl.R))
	}); dst != "" {
		gate.XYxor = dst
		if index == 0 {
			gate.Z = dst
		}
	}

	if dst := r.Match(func(gl GateLogic) bool {
		return gl.Op == OpAND && ((x == gl.L && y == gl.R) || (y == gl.L && x == gl.R))
	}); dst != "" {
		gate.XYand = dst

		if index == 0 {
			gate.Cout = dst
		}
	}

	if dst := r.Match(func(gl GateLogic) bool {
		return gl.Op == OpAND && (gl.L == gate.XYxor || gl.R == gate.XYxor)
	}); dst != "" {
		gate.Cand = dst

		if temp := r.gates[dst]; temp.L == gate.XYxor {
			gate.Cin = temp.R
		} else {
			gate.Cin = temp.L
		}
	}

	if dst := r.Match(func(gl GateLogic) bool {
		return gl.Op == OpOR && ((gl.L == gate.XYand && gl.R == gate.Cand) || (gl.R == gate.XYand && gl.L == gate.Cand))
	}); dst != "" {
		gate.Cout = dst
	}

	if dst := r.Match(func(gl GateLogic) bool {
		return gl.Op == OpXOR && ((gl.L == gate.XYxor && gl.R == gate.Cin) || (gl.R == gate.XYxor && gl.L == gate.Cin))
	}); dst != "" {
		gate.Z = dst
	}

	return &gate
}

func FindBrokenGates(resolver *Resolver) (swapped []string) {
	adders := []*Adder{}
	for index := range len(resolver.registers) / 2 {
		gate := MakeGate(resolver, index)
		adders = append(adders, gate)
	}

	// detect swapped gates
	for index, gate := range adders {
		if gate.Z != fmt.Sprintf("z%02d", index) {
			swapped = append(swapped, AnalyzeGate(resolver, adders, index)...)
		}
	}

	slices.Sort(swapped)
	return
}

func AnalyzeGate(r *Resolver, adders []*Adder, index int) (out []string) {
	brokenAdder := adders[index]

	// simple case: another field took z place:
	// return its id; z id easily inferred
	if len(brokenAdder.Z) > 0 {
		return []string{brokenAdder.Z, fmt.Sprintf("z%02d", index)}
	}

	// hard case: go backwards from known fields
	// comparing what we got with what we have.
	// The only fields that may be swapped inside
	// the adder are: xyXor, xyAnd, cAnd
	brokenAdder.Z = fmt.
		Sprintf("z%02d", index)
	// pull cin, cout from surrounding adders
	if index < len(adders)-1 {
		brokenAdder.Cout = adders[index+1].Cin
	}
	if index > 0 {
		brokenAdder.Cin = adders[index-1].Cout
	}

	var pivotGate GateLogic

	// check xyXor: find z, exclude cIn, the other input must be xyXor
	pivotGate = r.gates[brokenAdder.Z]
	xyXor := pivotGate.R
	if xyXor == brokenAdder.Cin {
		xyXor = pivotGate.L
	}
	if xyXor != brokenAdder.XYxor {
		return []string{xyXor, brokenAdder.XYxor}
	}

	// check cAnd, check xyAnd: not needed for this specific case
	return
}
