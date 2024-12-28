package lib

import (
	"fmt"
	"iter"
	"slices"
)

type Coords struct {
	X, Y int
}

type WithCoords[T any] struct {
	Value T
	Coords
}

func NewGrid[T any](w, h int, data []T) *Grid[T] {
	return &Grid[T]{
		items: data,
		w:     w,
		h:     h,
	}
}

type Grid[T any] struct {
	items []T
	w, h  int
}

func (g *Grid[T]) Width() int {
	return g.w
}

func (g *Grid[T]) Height() int {
	return g.h
}

func (g *Grid[T]) Clone() *Grid[T] {
	return NewGrid(g.w, g.h, slices.Clone(g.items))
}

func (g *Grid[T]) Find(predicate func(WithCoords[T]) bool) (WithCoords[T], bool) {
	for _, item := range g.Iter() {
		if predicate(item) {
			return item, true
		}
	}

	return WithCoords[T]{}, false
}

func (g *Grid[T]) FindAll(predicate func(WithCoords[T]) bool) iter.Seq[WithCoords[T]] {
	return func(yield func(WithCoords[T]) bool) {
		for _, item := range g.Iter() {
			if predicate(item) {
				if ok := yield(item); !ok {
					return
				}
			}
		}
	}
}

func (g *Grid[T]) index(x, y int) int {
	return g.w*y + x
}

func (g *Grid[T]) coords(index int) (x, y int) {
	return index % g.w, index / g.w
}

func (g *Grid[T]) At(x, y int) (value T, ok bool) {
	if x < 0 || y < 0 || x >= g.w || y >= g.h {
		return
	}

	index := g.index(x, y)

	if index < 0 || index >= len(g.items) {
		return
	}

	return g.items[index], true
}

func (g *Grid[T]) Set(x, y int, value T) {
	g.items[g.index(x, y)] = value
}

func (g *Grid[T]) AtUnsafe(x, y int) T {
	c, _ := g.At(x, y)
	return c
}

func (g *Grid[T]) Iter() iter.Seq2[int, WithCoords[T]] {
	return func(yield func(int, WithCoords[T]) bool) {
		for i, value := range g.items {
			x, y := g.coords(i)
			item := WithCoords[T]{value, Coords{x, y}}
			if ok := yield(i, item); !ok {
				return
			}
		}
	}
}

func Deltas8(x, y int) [][2]int {
	return [][2]int{
		{x - 1, y - 1}, {x, y - 1}, {x + 1, y - 1},
		{x - 1, y} /* current  tile */, {x + 1, y},
		{x - 1, y + 1}, {x, y + 1}, {x + 1, y + 1},
	}
}

func Deltas4(x, y int) [][2]int {
	return [][2]int{{x, y - 1}, {x + 1, y}, {x, y + 1}, {x - 1, y}}
}

func (g *Grid[T]) Around(x, y int, deltas func(x, y int) [][2]int) iter.Seq[WithCoords[T]] {
	return func(yield func(WithCoords[T]) bool) {
		for _, tuple := range deltas(x, y) {
			x, y := tuple[0], tuple[1]
			value, ok := g.At(x, y)

			if !ok {
				continue
			}

			if ok := yield(WithCoords[T]{value, Coords{x, y}}); !ok {
				return
			}
		}
	}
}

func (g *Grid[T]) Print(charfunc func(WithCoords[T]) rune) {
	for index, item := range g.Iter() {
		fmt.Print(string(charfunc(item)))
		if index > 0 && (index+1)%g.w == 0 {
			fmt.Println()
		}
	}
	fmt.Println()
}
