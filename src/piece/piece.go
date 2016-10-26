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

func (piece Piece) CanMove(diff point.Point) bool {
	canMove := false
	for i := 0; i < len(piece.movable); i++ {
		if diff.Y == piece.movable[i].Y && diff.X == piece.movable[i].X {
			canMove = true
		}
	}
	return canMove
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
