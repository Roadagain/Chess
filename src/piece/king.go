package piece

import "point"

var (
	kingMovable      []point.Point
	kingFirstMovable []point.Point
	King             Piece
)

func init() {
	for _, i := range [3]int{-1, 0, 1} {
		for _, j := range [3]int{-1, 0, 1} {
			if i == 0 && j == 0 {
				continue
			}
			kingMovable = append(kingMovable, point.Point{i, j})
		}
	}
	kingFirstMovable = append(kingFirstMovable, point.Point{0, 2})
	kingFirstMovable = append(kingFirstMovable, point.Point{0, -2})
	King = Piece{
		movable:      kingMovable,
		firstMovable: kingFirstMovable,
		white:        'K',
		black:        'k',
	}
}
