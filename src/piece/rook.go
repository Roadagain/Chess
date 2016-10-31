package piece

import "matrix"

var (
	rookMovable []matrix.Point
	Rook        Piece
)

func init() {
	for i := 1; i < matrix.SIDE; i++ {
		rookMovable = append(rookMovable, matrix.Point{i, 0})
		rookMovable = append(rookMovable, matrix.Point{0, i})
		rookMovable = append(rookMovable, matrix.Point{-i, 0})
		rookMovable = append(rookMovable, matrix.Point{0, -i})
	}
	Rook = Piece{
		movable:      rookMovable,
		firstMovable: make([]matrix.Point, 0),
		white:        'R',
		black:        'r',
	}
}
