package day7

import (
	"iter"
	"slices"
)

func Permute[T any](alphabet []T, len int) iter.Seq[[]T] {
	return permuteRec(nil, alphabet, len)
}

func permuteRec[T any](base []T, variants []T, remaining int) iter.Seq[[]T] {
	if remaining == 0 {
		return func(yield func([]T) bool) { yield(base) }
	}

	var zero T

	return func(yield func([]T) bool) {
		vec := append(slices.Clone(base), zero)
		for _, v := range variants {
			vec[len(vec)-1] = v
			for result := range permuteRec(vec, variants, remaining-1) {
				if ok := yield(result); !ok {
					return
				}
			}
		}

	}
}

type Op rune

const (
	OpAdd Op = '+'
	OpMul Op = '×'
	OpCat Op = '|'
)

func Valid(expected int, ops []Op, operands ...int) bool {
	for ops := range Permute(ops, len(operands)-1) {
		result := operands[0]
		for index, op := range ops {
			result = applyOp(op, result, operands[index+1])
		}
		if result == expected {
			return true
		}
	}

	return false
}

func applyOp(op Op, left, right int) int {
	switch op {
	case OpAdd:
		return left + right
	case OpMul:
		return left * right
	case OpCat:
		return catint(left, right)
	}
	panic("unreachable")
}

func catint(a, b int) int {
	tmp := b

	for tmp > 0 {
		a *= 10
		tmp /= 10
	}

	return a + b
}
