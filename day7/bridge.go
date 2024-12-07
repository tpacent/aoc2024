package day7

import (
	"iter"
)

type Task struct {
	Result int
	Nums   []int
}

func MatchExpr(task Task, ops []Op) bool {
	for value := range permuteNums(task.Result, ops, task.Nums[1:], task.Nums[0]) {
		if value == task.Result {
			return true
		}
	}

	return false
}

func permuteNums(target int, ops []Op, operands []int, prev int) iter.Seq[int] {
	if prev > target {
		return func(yield func(int) bool) {}
	}

	if len(operands) == 1 {
		return func(yield func(int) bool) {
			for _, op := range ops {
				if ok := yield(applyOp(op, prev, operands[0])); !ok {
					return
				}
			}
		}
	}

	return func(yield func(int) bool) {
		for _, op := range ops {
			for result := range permuteNums(target, ops, operands[1:], applyOp(op, prev, operands[0])) {
				if ok := yield(result); !ok {
					return
				}
			}
		}
	}
}

type Op byte

const (
	OpAdd Op = '+'
	OpMul Op = '×'
	OpCat Op = '|'
)

func applyOp(op Op, left, right int) int {
	switch op {
	case OpAdd:
		return left + right
	case OpMul:
		return left * right
	case OpCat:
		return catint(left, right)
	}
	return 0
}

func catint(a, b int) int {
	tmp := b

	for tmp > 0 {
		a *= 10
		tmp /= 10
	}

	return a + b
}
