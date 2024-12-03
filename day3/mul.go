package day3

import (
	"aoc24/lib"
	"regexp"
	"strings"
)

var ReMul = regexp.MustCompile(`(?m)mul\((\d+),(\d+)\)`)

func MulSum(s string) (total int) {
	for _, match := range ReMul.FindAllStringSubmatch(s, -1) {
		total += lib.MustParse(match[1]) * lib.MustParse(match[2])
	}

	return
}

var ReMulToggle = regexp.MustCompile(`(?m)mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)

func MulSumToggle(s string) (total int) {
	enabled := true

	for _, match := range ReMulToggle.FindAllStringSubmatch(s, -1) {
		if strings.HasPrefix(match[0], "mul") {
			if enabled {
				total += lib.MustParse(match[1]) * lib.MustParse(match[2])
			}
			continue
		}

		if strings.HasPrefix(match[0], "don") {
			enabled = false
			continue
		}

		if strings.HasPrefix(match[0], "do") {
			enabled = true
			continue
		}
	}

	return
}
