package lib

import (
	"strconv"
)

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
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
