package piece

import "point"

type Piece struct {
	movable []point.Point
	white   byte
	black   byte
}

func NewPiece(movable []point.Point, white, black byte) *Piece {
	piece := new(Piece)
	piece.movable = movable
	piece.white = white
	piece.black = black
	return piece
}

func (piece Piece) CanMove(diff point.Point) bool {
	for i := 0; i < len(piece.movable); i++ {
		if diff.Y == piece.movable[i].Y && diff.X == piece.movable[i].X {
			return true
		}
	}
	return false
}
