package enemy

import (
	"board"
	"color"
)

type Enemy struct {
	board.Board
	color.Color
	first [8][8]bool
}

func NewEnemy(b board.Board, c color.Color, first [8][8]bool) *Enemy {
	enemy := new(Enemy)
	enemy.Board = b
	enemy.Color = c
	enemy.first = first
	return enemy
}
