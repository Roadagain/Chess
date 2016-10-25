package piece

import "point"

var (
	pawnMovable []point.Point
	Pawn        Piece
)

func init() {
	pawnMovable.append(point.Point{-1, 0})
	Pawn = Piece{
		movable: pawnMovable,
		white:   'P',
		black:   'p',
	}
}
