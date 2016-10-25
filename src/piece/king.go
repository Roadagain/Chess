package piece

import "point"

var (
	kingMove []point.Point
	King     Piece
)

func init() {
	for _, i := range [3]int{-1, 0, 1} {
		for _, j := range [3]int{-1, 0, 1} {
			if i == 0 && j == 0 {
				continue
			}
			kingMove.append(point.Point{i, j})
		}
	}
	King = Piece{
		movable: kingMove,
		white:   'K',
		black:   'k',
	}
}
