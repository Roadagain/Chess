package piece

import "matrix"

var Empty = Piece{
	movable:      make([]matrix.Point, 0),
	firstMovable: make([]matrix.Point, 0),
	white:        ' ',
	black:        ' ',
}
