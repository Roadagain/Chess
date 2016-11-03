package piece

import (
	"color"
	"matrix"
)

type Piece struct {
	movable      []matrix.Point
	firstMovable []matrix.Point
	enemyMovable []matrix.Point
	white        byte
	black        byte
}

func NewPiece(movable, firstMovable, enemyMovable []matrix.Point, white, black byte) *Piece {
	piece := new(Piece)
	piece.movable = movable
	piece.firstMovable = firstMovable
	piece.enemyMovable = enemyMovable
	piece.white = white
	piece.black = black
	return piece
}

func (piece Piece) CanMove(diff matrix.Point, first, enemy bool) bool {
	for _, i := range piece.movable {
		if diff.Y == i.Y && diff.X == i.X && (Pawn.IsSymbol(piece.white) == false || enemy == false) {
			return true
		}
	}
	for _, i := range piece.firstMovable {
		if diff.Y == i.Y && diff.X == i.X && first && (Pawn.IsSymbol(piece.white) == false || enemy == false) {
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

func (piece Piece) Symbol(c color.Color) byte {
	if c == color.White {
		return piece.white
	} else if c == color.Black {
		return piece.black
	} else {
		return ' '
	}
}
