package piece

import "point"

var (
	rookMove []point.Point
	Rook     Piece
)

func init() {
	for i = 1; i < 8; i++ {
		rookMove.append(point.Point{i, 0})
		rookMove.append(point.Point{0, i})
		rookMove.append(point.Point{-i, 0})
		rookMove.append(point.Point{0, -i})
	}
	Rook = Piece{
		movable: rookMove,
		white:   'R',
		black:   'r',
	}
}
