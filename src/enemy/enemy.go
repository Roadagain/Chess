package enemy

import (
	"board"
	"color"
	"matrix"
)

type Enemy struct {
	matrix.Matrix
	color.Color
	first [8][8]bool
}

func NewEnemy(b *board.Board, c color.Color) *Enemy {
	enemy := new(Enemy)
	enemy.Matrix = b.Matrix
	enemy.Color = c
	for i := 0; i < matrix.SIDE; i++ {
		for j := 0; j < matrix.SIDE; j++ {
			enemy.first[i][j] = b.IsFirst(matrix.Point{i, j})
		}
	}
	return enemy
}
