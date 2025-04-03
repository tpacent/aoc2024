package lib

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func JoinInts[T Integer](values []T) string {
	ss := make([]string, 0, len(values))
	for _, v := range values {
		ss = append(ss, strconv.Itoa(int(v)))
	}
	return strings.Join(ss, ",")
}

func PrintResult[T comparable](t *testing.T, day, part int, actual, expected T) {
	header := fmt.Sprintf("Day %2d (%d):", day, part)

	if actual != expected {
		t.Error(header, "unexpected value")
		return
	}

	t.Logf("%-16s %v\n", header, actual)
}
