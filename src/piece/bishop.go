package piece

import "point"

var (
	bishopMove []point.Point
	Bishop     Piece
)

func init() {
	for i := 1; i < 8; i++ {
		bishopMove.append(point.Point{i, i})
		bishopMove.append(point.Point{-i, i})
		bishopMove.append(point.Point{i, -i})
		bishopMove.append(point.Point{-i, -i})
	}
	Bishop = Piece{
		movable: bishopMove,
		white:   'B',
		black:   'b',
	}
}
