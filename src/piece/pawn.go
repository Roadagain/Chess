package piece

import "point"

var (
	pawnMove []point.Point
	Pawn     Piece
)

func init() {
	pawnMove.append(point.Point{-1, 0})
	Pawn = Piece{
		movable: pawnMove,
		white:   'P',
		black:   'p',
	}
}
