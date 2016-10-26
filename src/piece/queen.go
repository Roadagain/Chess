package piece

import "point"

var (
	queenMovable []point.Point
	Queen        Piece
)

func init() {
	for i := 1; i < 8; i++ {
		queenMovable = append(queenMovable, point.Point{-i, -i})
		queenMovable = append(queenMovable, point.Point{-i, 0})
		queenMovable = append(queenMovable, point.Point{-i, i})
		queenMovable = append(queenMovable, point.Point{0, -i})
		queenMovable = append(queenMovable, point.Point{0, i})
		queenMovable = append(queenMovable, point.Point{i, -i})
		queenMovable = append(queenMovable, point.Point{i, 0})
		queenMovable = append(queenMovable, point.Point{i, i})
	}
	Queen = Piece{
		movable:      queenMovable,
		firstMovable: make([]point.Point, 0),
		white:        'Q',
		black:        'q',
	}
}
