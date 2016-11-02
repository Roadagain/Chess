package enemy

import (
	"color"
	"math/rand"
	"matrix"
	"piece"
	"time"
)

func (enemy Enemy) RandomizedSelect() (matrix.Point, matrix.Point) {
	rand.Seed(time.Now().UnixNano())

	var from, to matrix.Point
	unsettled := true
	mine := enemy.Matrix.Positions(enemy.Color)
	movable := append(enemy.Matrix.Positions(color.Empty), enemy.Matrix.Positions(enemy.Color.Enemy())...)

	for _, i := range mine {
		p := piece.WhichPiece(enemy.Matrix[i.Y][i.X])
		for _, j := range movable {
			var diff matrix.Point
			if enemy.Color == color.White {
				diff = i.Diff(j)
			} else {
				diff = j.Diff(i)
			}

			toEnemy := color.WhichColor(enemy.Matrix[j.Y][j.X]) == enemy.Color.Enemy()
			if p.CanMove(diff, enemy.first[i.Y][i.X], toEnemy) && enemy.Matrix.ExistBarrier(i, j) == false {
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
