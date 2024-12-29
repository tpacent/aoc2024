package lib

import (
	"strconv"
	"strings"
)

func JoinInts[T Integer](values []T) string {
	ss := make([]string, 0, len(values))
	for _, v := range values {
		ss = append(ss, strconv.Itoa(int(v)))
	}
	return strings.Join(ss, ",")
}
