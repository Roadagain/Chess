package piece

import "point"

var kingMove = []point.Point{
	point.Point{-1, -1},
	point.Point{-1, 0},
	point.Point{-1, 1},
	point.Point{0, -1},
	point.Point{0, 1},
	point.Point{1, -1},
	point.Point{1, 0},
	point.Point{1, 1},
}
var King = NewPiece(kingMove)
