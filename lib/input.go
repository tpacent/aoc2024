package lib

import (
	"bufio"
	"io"
	"iter"
)

func ReadInput[T any](r io.Reader, f func(string) T) iter.Seq[T] {
	scanner := bufio.NewScanner(r)

	return func(yield func(T) bool) {
		for scanner.Scan() {
			val := f(scanner.Text())
			if ok := yield(val); !ok {
				return
			}
		}
	}
}
