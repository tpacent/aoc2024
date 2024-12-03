package day3

import (
	"aoc24/lib"
	"bufio"
	"errors"
	"fmt"
	"io"
	"iter"
	"strconv"
	"strings"
)

var softError = errors.New("softfail")

type cmdSpec struct {
	Len int
}

var commands = map[string]cmdSpec{
	"do":    {Len: 0},
	"mul":   {Len: 2},
	"don't": {Len: 0},
}

type Cmd struct {
	Name  string
	Parms []int
}

func ParseIter(src io.Reader) iter.Seq[Cmd] {
	r := bufio.NewReader(src)

	return func(yield func(Cmd) bool) {
		for {
			name, err := ScanCommand(r)
			if err != nil {
				if errors.Is(err, softError) {
					continue
				} else {
					return
				}
			}

			parms, err := ScanParams(r)
			if err != nil {
				if errors.Is(err, softError) {
					continue
				} else {
					return
				}
			}

			spec := commands[name]

			if len(parms) != spec.Len {
				continue
			}

			cmd := Cmd{
				Name:  name,
				Parms: parms,
			}

			if ok := yield(cmd); !ok {
				return
			}
		}
	}
}

const (
	TokParmBegin = '('
	TokParmEnd   = ')'
	TokParmSep   = ','
)

func ScanCommand(r *bufio.Reader) (_ string, err error) {
	var (
		next []byte
		b    byte
		acc  []byte
	)

	for {
		if b, err = r.ReadByte(); err != nil {
			return "", err
		}

		next, err = r.Peek(1)
		if err != nil {
			return "", err
		}

		isCommandDone := next[0] == TokParmBegin
		acc = append(acc, b)
		chunk := string(acc)
		discard := true

		if isCommandDone {
			if _, ok := commands[chunk]; ok {
				return chunk, nil
			}

			acc = nil
			continue
		}

		for cmd := range commands {
			if strings.HasPrefix(cmd, chunk) {
				discard = false
				break
			}
		}

		if !discard {
			continue
		}

		// no command matched; clear buffer
		acc = acc[:0]
	}
}

func ScanParams(r *bufio.Reader) (params []int, err error) {
	var c byte
	var acc []byte

	if c, err = r.ReadByte(); err != nil {
		return nil, err
	}
	if c != TokParmBegin {
		return nil, fmt.Errorf("%w: unexpected start token", softError)
	}

	for {
		if c, err = r.ReadByte(); err != nil {
			return nil, err
		}

		if c == TokParmSep {
			if len(acc) > 0 {
				params = append(params, lib.MustParse(string(acc)))
			} else {
				return nil, fmt.Errorf("%w: unexpected separator", softError)
			}

			acc = nil
			continue
		}

		if c == TokParmEnd {
			if len(acc) > 0 {
				n, _ := strconv.Atoi(string(acc))
				params = append(params, n)
				acc = nil
			}

			return
		}

		if '0' <= c && '9' >= c {
			acc = append(acc, c)
			continue
		}

		return nil, fmt.Errorf("%w: unexpected token", softError)
	}
}
