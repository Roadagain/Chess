package piece

import "point"

var queenMove = []point.Point{
	point.Point{-8, -8},
	point.Point{-8, 8},
	point.Point{8, -8},
	point.Point{8, 8},
}
var Queen = NewPiece(queenMove)
