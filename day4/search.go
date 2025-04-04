package day4

import (
	"aoc24/lib"
	"bytes"
)

var directions = []func(int, int) (int, int){
	func(x, y int) (int, int) { return x - 1, y - 1 }, // 🡔
	func(x, y int) (int, int) { return x + 0, y - 1 }, // 🡑
	func(x, y int) (int, int) { return x + 1, y - 1 }, // 🡕
	func(x, y int) (int, int) { return x - 1, y + 0 }, // 🡐
	func(x, y int) (int, int) { return x + 1, y + 0 }, // 🡒
	func(x, y int) (int, int) { return x - 1, y + 1 }, // 🡗
	func(x, y int) (int, int) { return x + 0, y + 1 }, // 🡓
	func(x, y int) (int, int) { return x + 1, y + 1 }, // 🡖
}

func SearchDir(grid *lib.Grid[byte], word []byte) (total int) {
	for _, item := range grid.Iter() {
		total += FindWordsDir(grid, item.X, item.Y, word)
	}

	return
}

func FindWordsDir(grid *lib.Grid[byte], x, y int, word []byte) (count int) {
	c, ok := grid.At(x, y)
	if !ok {
		return
	}

	if c == word[0] {
		for _, next := range directions {
			if nx, ny := next(x, y); walkdir(grid, nx, ny, word[1:], next) {
				count++
			}
		}
	}

	return
}

func walkdir(grid *lib.Grid[byte], x, y int, word []byte, next func(int, int) (int, int)) bool {
	for _, expected := range word {
		c, ok := grid.At(x, y)

		if !ok {
			return false
		}

		if c != expected {
			return false
		}

		x, y = next(x, y)
	}

	return true
}

func SearchXMas(grid *lib.Grid[byte]) (total int) {
	for _, item := range grid.Iter() {
		if checkXMAS(grid, item.X, item.Y) {
			total++
		}
	}

	return
}

func checkXMAS(grid *lib.Grid[byte], x, y int) bool {
	pivot, ok := grid.At(x, y)

	if !ok {
		return false
	}

	if pivot != 'A' {
		return false
	}

	corners := make([]byte, 4, 8)
	corners[0] = grid.AtUnsafe(x-1, y-1)
	corners[1] = grid.AtUnsafe(x+1, y-1)
	corners[2] = grid.AtUnsafe(x+1, y+1)
	corners[3] = grid.AtUnsafe(x-1, y+1)

	template := []byte("MMSS")

	// rotations
	for range corners {
		if bytes.Equal(corners, template) {
			return true
		}

		corners = append(corners[1:], corners[0])
	}

	return false
}
