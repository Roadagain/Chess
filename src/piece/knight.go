package piece

import "point"

var (
	knightMove []point.Point
	Knight     Piece
)

func init() {
	for _, i := range [2]int{-1, 1} {
		for _, j := range [2]int{-2, 2} {
			knightMove.append(point.Point{i, j})
			knightMove.append(point.Point{j, i})
		}
	}
	Knight = Piece{
		movable: knightMove,
		white:   'N',
		black:   'n',
	}
}
