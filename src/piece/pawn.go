package piece

import "point"

var (
	pawnMovable      []point.Point
	pawnFirstMovable []point.Point
	pawnEnemyMovable []point.Point
	Pawn             Piece
)

func init() {
	pawnMovable = append(pawnMovable, point.Point{-1, 0})
	pawnFirstMovable = append(pawnFirstMovable, point.Point{-2, 0})
	pawnEnemyMovable = append(pawnEnemyMovable, point.Point{-1, -1})
	pawnEnemyMovable = append(pawnEnemyMovable, point.Point{-1, 1})
	Pawn = Piece{
		movable:      pawnMovable,
		firstMovable: pawnFirstMovable,
		enemyMovable: pawnEnemyMovable,
		white:        'P',
		black:        'p',
	}
}
