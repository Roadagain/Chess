package matrix

import "point"

type Matrix [8][8]byte

var starting = Matrix{
	{'r', 'n', 'b', 'k', 'q', 'b', 'n', 'r'},
	{'p', 'p', 'p', 'p', 'p', 'p', 'p', 'p'},
	{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
	{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
	{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
	{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
	{'P', 'P', 'P', 'P', 'P', 'P', 'P', 'P'},
	{'R', 'N', 'B', 'K', 'Q', 'B', 'N', 'R'},
}

func Starting() Matrix {
	return starting
}

func InMatrix(p point.Point) bool {
	return 0 <= p.X && p.X < 8 && 0 <= p.Y && p.Y < 8
}

func (mat Matrix) ExistBarrier(from, to point.Point) bool {
	p := from
	for p != to {
		if p.Y < to.Y {
			p.Y++
		} else if p.Y > to.Y {
			p.Y--
		}
		if p.X < to.X {
			p.X++
		} else if p.X > to.X {
			p.X--
		}
		if mat[p.Y][p.X] != ' ' {
			return true
		}
	}
	return false
}
