package piece

import "matrix"

var (
	queenMovable []matrix.Point
	Queen        Piece
)

func init() {
	for i := 1; i < matrix.SIDE; i++ {
		queenMovable = append(queenMovable, matrix.Point{-i, -i})
		queenMovable = append(queenMovable, matrix.Point{-i, 0})
		queenMovable = append(queenMovable, matrix.Point{-i, i})
		queenMovable = append(queenMovable, matrix.Point{0, -i})
		queenMovable = append(queenMovable, matrix.Point{0, i})
		queenMovable = append(queenMovable, matrix.Point{i, -i})
		queenMovable = append(queenMovable, matrix.Point{i, 0})
		queenMovable = append(queenMovable, matrix.Point{i, i})
	}
	Queen = Piece{
		movable: queenMovable,
		white:   'Q',
		black:   'q',
		score:   9,
	}
}
