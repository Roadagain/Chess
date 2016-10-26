package piece

import "point"

var (
	rookMovable []point.Point
	Rook        Piece
)

func init() {
	for i := 1; i < 8; i++ {
		rookMovable = append(rookMovable, point.Point{i, 0})
		rookMovable = append(rookMovable, point.Point{0, i})
		rookMovable = append(rookMovable, point.Point{-i, 0})
		rookMovable = append(rookMovable, point.Point{0, -i})
	}
	Rook = Piece{
		movable:      rookMovable,
		firstMovable: make([]point.Point, 0),
		white:        'R',
		black:        'r',
	}
}