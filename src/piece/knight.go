package piece

import "point"

var knightMove = []point.Point{
	point.Point{-2, -1},
	point.Point{-2, 1},
	point.Point{-1, -2},
	point.Point{-1, 2},
	point.Point{1, -2},
	point.Point{1, 2},
	point.Point{2, -1},
	point.Point{2, 1},
}
var Knight = NewPiece(knightMove, 'N', 'n')
