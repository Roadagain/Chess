package piece

import "point"

var pawnMove = []point.Point{
	point.Point{-1, 0},
}
var Pawn = Piece{
	movable: pawnMove,
	white:   'P',
	black:   'p',
}
