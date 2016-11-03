package matrix

import (
	"color"
	"fmt"
)

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

func (mat Matrix) ToString() string {
	matBytes := make([]byte, 0)
	matBytes = append(matBytes, " "...)
	//matBytes := " "
	for i := 0; i < SIDE; i++ {
		matBytes = append(matBytes, fmt.Sprintf(" %c", 'a'+i)...)
	}
	matBytes = append(matBytes, "\n"...)

	for i := 0; i < SIDE; i++ {
		matBytes = append(matBytes, " +"...)
		for i := 0; i < SIDE; i++ {
			matBytes = append(matBytes, "-+"...)
		}
		matBytes = append(matBytes, "\n"...)
		matBytes = append(matBytes, fmt.Sprintf("%d|", SIDE-i)...)
		for j := 0; j < SIDE; j++ {
			matBytes = append(matBytes, fmt.Sprintf("%c|", mat[i][j])...)
		}
		matBytes = append(matBytes, fmt.Sprintf("%d\n", SIDE-i)...)
	}
	matBytes = append(matBytes, " +"...)
	for i := 0; i < SIDE; i++ {
		matBytes = append(matBytes, "-+"...)
	}
	matBytes = append(matBytes, "\n"...)

	matBytes = append(matBytes, " "...)
	for i := 0; i < SIDE; i++ {
		matBytes = append(matBytes, fmt.Sprintf(" %c", 'a'+i)...)
	}

	return string(matBytes)
}

func (mat Matrix) Positions(c color.Color) []Point {
	positions := make([]Point, 0)
	for i := 0; i < SIDE; i++ {
		for j := 0; j < SIDE; j++ {
			if color.WhichColor(mat[i][j]) == c {
				positions = append(positions, Point{i, j})
			}
		}
	}

	return positions
}
