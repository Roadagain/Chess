package piece

import "point"

var (
	bishopMovable []point.Point
	Bishop        Piece
)

func init() {
	for i := 1; i < 8; i++ {
		bishopMovable = append(bishopMovable, point.Point{i, i})
		bishopMovable = append(bishopMovable, point.Point{-i, i})
		bishopMovable = append(bishopMovable, point.Point{i, -i})
		bishopMovable = append(bishopMovable, point.Point{-i, -i})
	}
	Bishop = Piece{
		movable:      bishopMovable,
		firstMovable: make([]point.Point, 0),
		white:        'B',
		black:        'b',
	}
}