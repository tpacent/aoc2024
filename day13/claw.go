package day13

import "iter"

func CramersRule(ax, ay, bx, by, tx, ty int) (int, int) {
	// Cramer’s Rule for 2×2 Systems
	// | a b |
	// | c d |
	// a * aSpec.Dx + b * bSpec.Dx == target.X
	// a * aSpec.Dy + b * bSpec.Dy == target.Y

	a1 := det(tx, bx, ty, by)
	a2 := det(ax, bx, ay, by)
	b1 := det(ax, tx, ay, ty)
	b2 := det(ax, bx, ay, by)
	a, b := a1/a2, b1/b2

	if a*ax+b*bx != tx || a*ay+b*by != ty {
		return 0, 0
	}

	return a, b
}

func det(a, b, c, d int) int {
	return a*d - b*c
}

type Button struct {
	Dx, Dy int
}

type Target struct {
	X, Y int
}

type Input struct {
	AButton Button
	BButton Button
	Target  Target
}

func CalcTokens(it iter.Seq[Input], offset int) (total int) {
	for task := range it {
		tx := offset + task.Target.X
		ty := offset + task.Target.Y
		ax := task.AButton.Dx
		ay := task.AButton.Dy
		bx := task.BButton.Dx
		by := task.BButton.Dy
		a, b := CramersRule(ax, ay, bx, by, tx, ty)
		total += a*3 + b
	}
	return
}
