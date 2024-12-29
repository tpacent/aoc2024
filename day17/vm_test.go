package day17_test

import (
	"aoc24/day17"
	"aoc24/lib"
	"fmt"
	"slices"
	"testing"
)

func TestDay17Part1(t *testing.T) {
	vm := day17.NewVM([]int{2, 4, 1, 2, 7, 5, 1, 3, 4, 3, 5, 5, 0, 3, 3, 0}, 64584136, 0, 0)
	vm.Run()
	t.Log(lib.JoinInts(vm.Out())) // 3,7,1,7,2,1,0,6,3
}

func TestDay17Part2(t *testing.T) {
	expected := []int{2, 4, 1, 2, 7, 5, 1, 3, 4, 3, 5, 5, 0, 3, 3, 0}
	values := day17.CodeBreaker(0, expected, 1)
	slices.Sort(values)
	t.Log(values[0]) // 37221334433268
}

func TestExample(t *testing.T) {
	tcases := []struct {
		Prg          []int
		A, B, C      int
		expectedOut  []int
		expectedRegs []int
	}{
		{C: 9, Prg: []int{2, 6}, expectedRegs: []int{0, 1, 0}},
		{A: 10, Prg: []int{5, 0, 5, 1, 5, 4}, expectedOut: []int{0, 1, 2}},
		{A: 2024, Prg: []int{0, 1, 5, 4, 3, 0}, expectedOut: []int{4, 2, 5, 6, 7, 7, 7, 7, 3, 1, 0}},
		{B: 29, Prg: []int{1, 7}, expectedRegs: []int{0, 26, 0}},
		{B: 2024, C: 43690, Prg: []int{4, 0}, expectedRegs: []int{0, 44354, 0}},
	}

	for index, tcase := range tcases {
		t.Run(fmt.Sprintf("%d", index), func(t *testing.T) {
			vm := day17.NewVM(tcase.Prg, tcase.A, tcase.B, tcase.C)
			vm.Run()

			if tcase.expectedOut != nil {
				if actual := vm.Out(); !slices.Equal(actual, tcase.expectedOut) {
					t.Errorf("unexpected output: %+v", actual)
				}
			}

			regs := []int{vm.A, vm.B, vm.C}
			for index, expected := range tcase.expectedRegs {
				if actual := regs[index]; expected != 0 && actual != expected {
					t.Errorf("unexpected register: got %d instead of %d", actual, expected)
				}
			}
		})
	}
}

func TestExample2(t *testing.T) {
	expected := []int{0, 3, 5, 4, 3, 0}
	values := day17.CodeBreaker(0, expected, 1)
	slices.Sort(values)
	if actual := values[0]; actual != 117440 {
		t.Error("unexpected value")
	}
}
