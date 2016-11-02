package enemy

import (
	"color"
	"math/rand"
	"matrix"
	"time"
)

func (enemy Enemy) RandomizedSelect() (matrix.Point, matrix.Point) {
	rand.Seed(time.Now().UnixNano())

	var from, to matrix.Point
	unsettled := true
	mine := enemy.Matrix.Positions(enemy.Color)
	movable := enemy.Board.Matrix.Positions(color.Empty)
	movable = append(movable, enemy.Board.Matrix.Positions(enemy.Color.Enemy())...)

	for _, i := range mine {
		for _, j := range movable {
			canMove, _ := enemy.Board.CanMove(i, j, enemy.Color)
			if canMove {
				if unsettled || rand.Intn(2) == 1 {
					from = i
					to = j
					unsettled = false
				}
			}
		}
	}

	return from, to
}
