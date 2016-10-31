package piece

import "matrix"

var (
	bishopMovable []matrix.Point
	Bishop        Piece
)

func init() {
	for i := 1; i < matrix.SIDE; i++ {
		bishopMovable = append(bishopMovable, matrix.Point{i, i})
		bishopMovable = append(bishopMovable, matrix.Point{-i, i})
		bishopMovable = append(bishopMovable, matrix.Point{i, -i})
		bishopMovable = append(bishopMovable, matrix.Point{-i, -i})
	}
	Bishop = Piece{
		movable:      bishopMovable,
		firstMovable: make([]matrix.Point, 0),
		white:        'B',
		black:        'b',
	}
}
