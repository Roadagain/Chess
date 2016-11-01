package matrix

const SIDE = 8

type Matrix [SIDE][SIDE]byte

var starting = Matrix{
	{'r', 'n', 'b', 'q', 'k', 'b', 'n', 'r'},
	{'p', 'p', 'p', 'p', 'p', 'p', 'p', 'p'},
	{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
	{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
	{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
	{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
	{'P', 'P', 'P', 'P', 'P', 'P', 'P', 'P'},
	{'R', 'N', 'B', 'Q', 'K', 'B', 'N', 'R'},
}

func Starting() Matrix {
	return starting
}

func InMatrix(p Point) bool {
	return 0 <= p.X && p.X < SIDE && 0 <= p.Y && p.Y < SIDE
}

func (mat Matrix) ExistBarrier(from, to Point) bool {
	p := from
	p.StepTo(to)
	for p != to {
		if mat[p.Y][p.X] != ' ' {
			return true
		}
		p.StepTo(to)
	}
	return false
}
