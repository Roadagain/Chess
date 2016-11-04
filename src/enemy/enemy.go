package enemy

import (
	"board"
	"color"
)

type Enemy struct {
	board.Board
	color.Color
}

const (
	Random = iota
	Brutal
)

func NewEnemy(b *board.Board, c color.Color) *Enemy {
	enemy := new(Enemy)
	enemy.Board = *b
	enemy.Color = c
	return enemy
}
