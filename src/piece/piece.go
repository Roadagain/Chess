package piece

import "point"

type Limit [3][3]int

type Piece struct {
	movable []point.Point
	limit   Limit
}
