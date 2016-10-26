package piece

import "point"

type Piece struct {
	movable      []point.Point
	firstMovable []point.Point
	white        byte
	black        byte
}

func NewPiece(movable, firstMovable []point.Point, white, black byte) *Piece {
	piece := new(Piece)
	piece.movable = movable
	piece.firstMovable = firstMovable
	piece.white = white
	piece.black = black
	return piece
}

func (piece Piece) CanMove(diff point.Point, first bool) bool {
	for i := 0; i < len(piece.movable); i++ {
		if diff.Y == piece.movable[i].Y && diff.X == piece.movable[i].X {
			return true
		}
	}
	for i := 0; i < len(piece.firstMovable); i++ {
		if diff.Y == piece.firstMovable[i].Y && diff.X == piece.firstMovable[i].X && first {
			return true
		}
	}
	return false
}

func (piece Piece) IsSymbol(symbol byte) bool {
	return symbol == piece.white || symbol == piece.black
}

func WhichPiece(symbol byte) Piece {
	switch symbol {
	case 'B', 'b':
		return Bishop
	case 'K', 'k':
		return King
	case 'N', 'n':
		return Knight
	case 'P', 'p':
		return Pawn
	case 'Q', 'q':
		return Queen
	case 'R', 'r':
		return Rook
	default:
		return Empty
	}
}
