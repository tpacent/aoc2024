package day3

import (
	"strings"
)

func MulSum(s string) (total int) {
	for cmd := range ParseIter(strings.NewReader(s)) {
		if cmd.Name == "mul" {
			total += cmd.Parms[0] * cmd.Parms[1]
		}
	}

	return
}

func MulSumToggle(s string) (total int) {
	mulEnabled := 1

	for cmd := range ParseIter(strings.NewReader(s)) {
		switch cmd.Name {
		case "mul":
			total += cmd.Parms[0] * cmd.Parms[1] * mulEnabled
		case "do":
			mulEnabled = 1
		case "don't":
			mulEnabled = 0
		}
	}

	return
}
