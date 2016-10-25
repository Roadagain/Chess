package piece

import "point"

var pawnMove = []point.Point{
	point.Point{-1, 0},
}
var Pawn = NewPiece(pawnMove, 'P', 'p')
