package piece

import "point"

var Empty = Piece{
	movable:      make([]point.Point, 0),
	firstMovable: make([]point.Point, 0),
	white:        ' ',
	black:        ' ',
}
