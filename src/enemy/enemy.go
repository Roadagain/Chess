package enemy

import (
	"board"
	"color"
	"strings"
)

type Enemy struct {
	board.Board
	color.Color
}

type Type int

const (
	Unknown = iota
	Random
	Brutal
)

func NewEnemy(b *board.Board, c color.Color) *Enemy {
	enemy := new(Enemy)
	enemy.Board = *b
	enemy.Color = c
	return enemy
}

func ParseType(s string) Type {
	switch strings.ToLower(s) {
	case "random":
		return Random
	case "brutal":
		return Brutal
	default:
		return Unknown
	}
}
