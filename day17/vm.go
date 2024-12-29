package day17

import (
	"slices"
)

type VM struct {
	instrp   int
	handlers []opHandler
	program  []int
	outbuf   []int
	// registers
	A int
	B int
	C int
}

func NewVM(program []int, a, b, c int) *VM {
	return &VM{
		program: program,
		A:       a,
		B:       b,
		C:       c,
		handlers: []opHandler{
			advHandler,
			bxlHandler,
			bstHandler,
			jnzHandler,
			bxcHandler,
			outHandler,
			bdvHandler,
			cdvHandler,
		},
	}
}

func (vm *VM) Out() []int {
	return vm.outbuf
}

func (vm *VM) Run() {
	for {
		if ok := vm.step(); !ok {
			break
		}
	}
}

func (vm *VM) step() bool {
	if vm.instrp < 0 || vm.instrp >= len(vm.program)-1 {
		return false
	}

	op := vm.program[vm.instrp]
	operand := vm.program[vm.instrp+1]
	vm.instrp += 2
	vm.handlers[op](vm, operand)
	return true
}

func (vm *VM) Combo(operand int) int {
	switch operand {
	case 0, 1, 2, 3:
		return operand
	case 4:
		return vm.A
	case 5:
		return vm.B
	case 6:
		return vm.C
	}

	panic("unreachable")
}

func (vm *VM) dv(a, operand int) int {
	return a / pow(2, vm.Combo(operand))
}

func pow(n, p int) int {
	if p == 0 {
		return 1
	}

	return n << (p - 1)
}

func CodeBreaker(initial int, expected []int, suffix int) (out []int) {
	for k := 0; k < 0o10; k++ {
		n := initial*0o10 + k
		vm := NewVM(expected, n, 0, 0)
		vm.Run()

		if vm.A > 0 || vm.B > 0 || vm.C > 0 {
			continue
		}

		actual := vm.Out()

		if len(actual) < suffix {
			continue
		}

		if slices.Equal(
			expected[len(expected)-suffix:],
			actual[len(actual)-suffix:],
		) {
			if suffix == len(expected) {
				out = append(out, n)
				continue
			}

			out = append(out, CodeBreaker(n, expected, suffix+1)...)
		}
	}

	return
}
