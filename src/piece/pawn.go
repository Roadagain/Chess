package piece

import "matrix"

var (
	pawnMovable      []matrix.Point
	pawnFirstMovable []matrix.Point
	pawnEnemyMovable []matrix.Point
	Pawn             Piece
)

func init() {
	pawnMovable = append(pawnMovable, matrix.Point{-1, 0})
	pawnFirstMovable = append(pawnFirstMovable, matrix.Point{-2, 0})
	pawnEnemyMovable = append(pawnEnemyMovable, matrix.Point{-1, -1})
	pawnEnemyMovable = append(pawnEnemyMovable, matrix.Point{-1, 1})
	Pawn = Piece{
		movable:      pawnMovable,
		firstMovable: pawnFirstMovable,
		enemyMovable: pawnEnemyMovable,
		white:        'P',
		black:        'p',
	}
}
