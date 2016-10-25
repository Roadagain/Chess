package piece

import "point"

var (
	knightMovable []point.Point
	Knight        Piece
)

func init() {
	for _, i := range [2]int{-1, 1} {
		for _, j := range [2]int{-2, 2} {
			knightMovable = append(knightMovable, point.Point{i, j})
			knightMovable = append(knightMovable, point.Point{j, i})
		}
	}
	Knight = Piece{
		movable: knightMovable,
		white:   'N',
		black:   'n',
	}
}
