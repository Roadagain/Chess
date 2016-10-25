package piece

import "point"

var (
	queenMove []point.Point
	Queen     Piece
)

func init() {
	for i = 1; i < 8; i++ {
		queenMove.append(point.Point{-i, -i})
		queenMove.append(point.Point{-i, 0})
		queenMove.append(point.Point{-i, i})
		queenMove.append(point.Point{0, -i})
		queenMove.append(point.Point{0, i})
		queenMove.append(point.Point{i, -i})
		queenMove.append(point.Point{i, 0})
		queenMove.append(point.Point{i, i})
	}
	Queen = Piece{
		movable: queenMove,
		white:   'Q',
		black:   'q',
	}
}
