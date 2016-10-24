package piece

import (
	"chessboard"
	"point"
)

type Piece struct {
	movable []point.Point
}

func NewPiece(movable []point.Point) *Piece {
	piece := new(Piece)
	piece.movable = movable
	return piece
}

func (piece Piece) CanMove(from, to point.Point) bool {
	if chessboard.InBoard(from) == false || chessboard.InBoard(to) == false {
		return false
	}
	diff := from.Diff(to)
	for i := 0; i < len(piece.movable); i++ {
		if diff.Y <= piece.movable[i].Y && diff.X <= piece.movable[i].X {
			return true
		}
	}
	return false
}
