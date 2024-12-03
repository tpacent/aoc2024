package lib

import (
	"strconv"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

func AbsDiff[T Number](a, b T) T {
	if a > b {
		return a - b
	}

	return b - a
}

func MustParse(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}
