package enemy

import (
	"color"
	"math/rand"
	"matrix"
	"piece"
	"time"
)

func (enemy Enemy) BrutalSelect() (matrix.Point, matrix.Point) {
	rand.Seed(time.Now().UnixNano())

	var from, to matrix.Point
	maxi := -1000
	mine := enemy.Matrix.Positions(enemy.Color)
	movable := enemy.Board.Matrix.Positions(color.Empty)
	movable = append(movable, enemy.Board.Matrix.Positions(enemy.Color.Enemy())...)

	for _, i := range mine {
		for _, j := range movable {
			canMove, _ := enemy.Board.CanMove(i, j, enemy.Color)
			if canMove {
				score := piece.WhichPiece(enemy.Matrix[j.Y][j.X]).Score()
				if maxi < score || (maxi == score && rand.Intn(2) == 1) {
					from = i
					to = j
					maxi = score
				}
			}
		}
	}

	return from, to
}
