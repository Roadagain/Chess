package piece

import "point"

var bishopMove = []point.Point{
	point.Point{-8, -8},
	point.Point{-8, 8},
	point.Point{8, -8},
	point.Point{8, 8},
}
var Bishop = NewPiece(bishopMove, 'B', 'b')
