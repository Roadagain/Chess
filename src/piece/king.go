package piece

import "matrix"

var (
	kingMovable      []matrix.Point
	kingFirstMovable []matrix.Point
	King             Piece
)

func init() {
	for _, i := range [3]int{-1, 0, 1} {
		for _, j := range [3]int{-1, 0, 1} {
			if i == 0 && j == 0 {
				continue
			}
			kingMovable = append(kingMovable, matrix.Point{i, j})
		}
	}
	kingFirstMovable = append(kingFirstMovable, matrix.Point{0, 2})
	kingFirstMovable = append(kingFirstMovable, matrix.Point{0, -2})
	King = Piece{
		movable:      kingMovable,
		firstMovable: kingFirstMovable,
		white:        'K',
		black:        'k',
	}
}
