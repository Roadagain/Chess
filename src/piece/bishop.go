package piece

import "point"

var bishopMove = []point.Point{
	point.Point{-8, -8},
	point.Point{-8, 8},
	point.Point{8, -8},
	point.Point{8, 8},
}
var Bishop = Piece{
	movable: bishopMove,
	white:   'B',
	black:   'b',
}
