package piece

import "point"

var (
	pawnMovable      []point.Point
	pawnFirstMovable []point.Point
	Pawn             Piece
)

func init() {
	pawnMovable = append(pawnMovable, point.Point{-1, 0})
	pawnFirstMovable = append(pawnFirstMovable, point.Point{-2, 0})
	Pawn = Piece{
		movable:      pawnMovable,
		firstMovable: pawnFirstMovable,
		white:        'P',
		black:        'p',
	}
}
