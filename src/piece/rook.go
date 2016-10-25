package piece

import "point"

var rookMove = []point.Point{
	point.Point{-8, 0},
	point.Point{0, -8},
	point.Point{0, 8},
	point.Point{8, 0},
}
var Rook = NewPiece(rookMove, 'R', 'r')
