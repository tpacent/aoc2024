package lib

import (
	"bufio"
	"bytes"
	"errors"
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

func ReadGrid[T any](src io.Reader, makeVal func(byte) T) (w, h int, data []T) {
	r := bufio.NewReader(src)

	for {
		line, err := r.ReadBytes('\n')
		chunk := bytes.TrimSuffix(line, []byte{'\n'})

		if len(chunk) > 0 {
			values := make([]T, 0, len(chunk))
			for _, b := range chunk {
				values = append(values, makeVal(b))
			}

			data = append(data, values...)
			h++
		}

		if errors.Is(err, io.EOF) {
			break
		}
	}

	w = len(data) / h
	return
}
