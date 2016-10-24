package chessboard

import "point"

type Move struct {
	file byte
	rank int
}

func NewMove(file byte, rank int) *Move {
	move := new(Move)
	move.file = file
	move.rank = rank
	return move
}

func (move Move) ToPoint() point.Point {
	return point.Point{8 - move.rank, int(move.file - 'a')}
}
