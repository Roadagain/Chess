package piece

import "matrix"

var (
	knightMovable []matrix.Point
	Knight        Piece
)

func init() {
	for _, i := range [2]int{-1, 1} {
		for _, j := range [2]int{-2, 2} {
			knightMovable = append(knightMovable, matrix.Point{i, j})
			knightMovable = append(knightMovable, matrix.Point{j, i})
		}
	}
	Knight = Piece{
		movable:      knightMovable,
		firstMovable: make([]matrix.Point, 0),
		white:        'N',
		black:        'n',
	}
}
