package lib

import (
	"bufio"
	"io"
	"iter"
	"os"
	"strconv"
	"strings"
)

func MustOpenFile(filename string) *os.File {
	if file, err := os.Open(filename); err != nil {
		panic(err)
	} else {
		return file
	}
}

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

func NumsLine(s string) (nums []int) {
	for _, f := range strings.Fields(s) {
		n, _ := strconv.Atoi(f)
		nums = append(nums, n)
	}
	return
}
