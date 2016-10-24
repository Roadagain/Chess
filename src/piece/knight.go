package piece

import "point"

var queenMove = []point.Point{
	point.Point{-2, -1},
	point.Point{-2, 1},
	point.Point{-1, -2},
	point.Point{-1, 2},
	point.Point{1, -2},
	point.Point{1, 2},
	point.Point{2, -1},
	point.Point{2, 1},
}
var Queen = NewPiece(queenMove)
