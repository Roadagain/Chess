package piece

import "point"

type Piece struct {
	movable      []point.Point
	firstMovable []point.Point
	enemyMovable []point.Point
	white        byte
	black        byte
}

func NewPiece(movable, firstMovable, enemyMovable []point.Point, white, black byte) *Piece {
	piece := new(Piece)
	piece.movable = movable
	piece.firstMovable = firstMovable
	piece.enemyMovable = enemyMovable
	piece.white = white
	piece.black = black
	return piece
}

func (piece Piece) CanMove(diff point.Point, first, enemy bool) bool {
	for _, i := range piece.movable {
		if diff.Y == i.Y && diff.X == i.X {
			return true
		}
	}
	for _, i := range piece.firstMovable {
		if diff.Y == i.Y && diff.X == i.X && first {
			return true
		}
	}
	for _, i := range piece.enemyMovable {
		if diff.Y == i.Y && diff.X == i.X && enemy {
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
