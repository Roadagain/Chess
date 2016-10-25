package piece

import "point"

var (
	kingMovable []point.Point
	King        Piece
)

func init() {
	for _, i := range [3]int{-1, 0, 1} {
		for _, j := range [3]int{-1, 0, 1} {
			if i == 0 && j == 0 {
				continue
			}
			kingMovable.append(point.Point{i, j})
		}
	}
	King = Piece{
		movable: kingMovable,
		white:   'K',
		black:   'k',
	}
}
