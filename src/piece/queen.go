package piece

import "point"

var (
	queenMovable []point.Point
	Queen        Piece
)

func init() {
	for i = 1; i < 8; i++ {
		queenMovable.append(point.Point{-i, -i})
		queenMovable.append(point.Point{-i, 0})
		queenMovable.append(point.Point{-i, i})
		queenMovable.append(point.Point{0, -i})
		queenMovable.append(point.Point{0, i})
		queenMovable.append(point.Point{i, -i})
		queenMovable.append(point.Point{i, 0})
		queenMovable.append(point.Point{i, i})
	}
	Queen = Piece{
		movable: queenMovable,
		white:   'Q',
		black:   'q',
	}
}
