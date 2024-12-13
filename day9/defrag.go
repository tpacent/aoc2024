package day9

import (
	"iter"
)

type Fragment struct {
	ID     int
	Size   int
	IsFile bool
}

func Checksum(iter iter.Seq[int]) (sum int) {
	index := 0
	for fileID := range iter {
		sum += index * fileID
		index++
	}
	return
}

func FileIter(input []Fragment) iter.Seq[int] {
	var (
		cursorL = 0
		cursorR = len(input) - 1
		rBuf    []int
	)

	return func(yield func(int) bool) {
		for {
			fragment := input[cursorL]

			if fragment.IsFile {
				for range fragment.Size {
					if ok := yield(fragment.ID); !ok {
						return
					}
				}
			} else {
				// fill rBuf to contain at least Size
				for len(rBuf) < fragment.Size {
					if cursorL >= cursorR {
						break
					}

					tail := input[cursorR]

					if !tail.IsFile {
						cursorR--
						continue
					}

					for range tail.Size {
						rBuf = append(rBuf, tail.ID)
					}

					input[cursorR].IsFile = false
					input[cursorR].ID = -1
					cursorR--
				}

				// consume from rbuf
				for range fragment.Size {
					if len(rBuf) == 0 {
						return
					}

					if ok := yield(rBuf[0]); !ok {
						return
					}
					rBuf = rBuf[1:]
				}
			}

			cursorL++
		}
	}
}

func FileIter2(input []Fragment) (cksum int) {
	reverseFiles := ReverseFiles(input)
	for _, file := range reverseFiles {

		// find hole
		holeIndex, hole := FindHole(input, file)

		if holeIndex < 0 {
			continue
		}

		remaining := hole.Size - file.Size

		input[holeIndex] = file

		if remaining > 0 {
			if !input[holeIndex+1].IsFile {
				// extend
				input[holeIndex+1].Size += remaining
			} else {
				input = append(input[:holeIndex+1], append([]Fragment{{Size: remaining, ID: -1}}, input[holeIndex+1:]...)...)
			}

		}

		ClearFile(input, file.ID)
	}

	index := 0
	for _, frag := range input {
		if !frag.IsFile {
			index += frag.Size
			continue
		}

		for range frag.Size {
			cksum += index * frag.ID
			index++
		}
	}

	return
}

func ClearFile(input []Fragment, id int) {
	var pivot int

	for pivot = len(input) - 1; pivot >= 0; pivot-- {
		if input[pivot].ID == id {
			break
		}
	}

	size := input[pivot].Size

	if pivot < len(input)-1 {
		if right := input[pivot+1]; !right.IsFile {
			size += right.Size
			input[pivot+1].Size = 0
		}
	}

	if pivot > 0 {
		if left := input[pivot-1]; !left.IsFile {
			size += left.Size
			input[pivot-1].Size = 0
		}
	}

	input[pivot] = Fragment{ID: -1, Size: size}
}

func FindHole(input []Fragment, src Fragment) (int, Fragment) {
	for index, frag := range input {
		if frag.IsFile {
			if frag.ID == src.ID {
				break
			}

			continue
		}

		if frag.Size < src.Size {
			continue
		}

		return index, frag
	}

	return -1, Fragment{}
}

func ReverseFiles(input []Fragment) []Fragment {
	reverse := make([]Fragment, 0, len(input))

	for k := len(input) - 1; k >= 0; k-- {
		if input[k].IsFile {
			reverse = append(reverse, input[k])
		}
	}

	return reverse
}
