package board

import (
	"color"
	"errors"
	"matrix"
	"piece"
)

type Board struct {
	matrix.Matrix
	first [8][8]bool
}

func NewBoard() *Board {
	board := new(Board)
	board.Matrix = matrix.Starting()
	for i := 0; i < 8; i++ {
		board.first[0][i] = true
		board.first[1][i] = true
		board.first[6][i] = true
		board.first[7][i] = true
	}
	return board
}

func (board Board) IsFirst(p matrix.Point) bool {
	return board.first[p.Y][p.X]
}

func (board *Board) CanCastling(from, to matrix.Point) bool {
	diff := from.Diff(to)
	rfrom := from
	rto := to

	if diff.X == 2 {
		rfrom.X = 7
		rto.X = to.X - 1
	} else {
		rfrom.X = 0
		rto.X = to.X + 1
	}
	fsymbol := board.Matrix[rfrom.Y][rfrom.X]
	if piece.Rook.IsSymbol(fsymbol) == false {
		return false
	} else if board.first[rfrom.Y][rfrom.X] == false {
		return false
	} else if board.Matrix.ExistBarrier(rfrom, rto) {
		return false
	}
	return true
}

func (board *Board) Castling(from, to matrix.Point) {
	diff := from.Diff(to)
	rfrom := from
	rto := to
	if diff.X == 2 {
		rfrom.X = 7
		rto.X = to.X - 1
	} else {
		rfrom.X = 0
		rto.X = to.X + 1
	}

	board.Matrix[to.Y][to.X] = board.Matrix[from.Y][from.X]
	board.Matrix[from.Y][from.X] = ' '
	board.first[from.Y][from.X] = false
	board.first[to.Y][to.X] = false
	board.Matrix[rto.Y][rto.X] = board.Matrix[rfrom.Y][rfrom.X]
	board.Matrix[rfrom.Y][rfrom.X] = ' '
	board.first[rfrom.Y][rfrom.X] = false
	board.first[rto.Y][rto.X] = false
}

func (board *Board) Move(from, to matrix.Point, c color.Color) error {
	if matrix.InMatrix(from) == false || matrix.InMatrix(to) == false {
		return errors.New("out of board")
	}

	fsymbol := board.Matrix[from.Y][from.X]
	tsymbol := board.Matrix[to.Y][to.X]

	fcolor := color.WhichColor(board.Matrix[from.Y][from.X])
	tcolor := color.WhichColor(board.Matrix[to.Y][to.X])
	diff := matrix.Point{0, 0}
	if fcolor == color.White {
		diff = from.Diff(to)
	} else if fcolor == color.Black {
		diff = to.Diff(from)
	}
	if fcolor != c || tcolor == c {
		return errors.New("cannot move pieces that is not yours")
	}

	fpiece := piece.WhichPiece(fsymbol)
	toEnemy := color.WhichColor(board.Matrix[to.Y][to.X]) == c.Enemy()
	canMove := fpiece.CanMove(diff, board.first[from.Y][from.X], toEnemy)
	existBarrier := board.Matrix.ExistBarrier(from, to)
	if canMove == false || (existBarrier && piece.Knight.IsSymbol(fsymbol) == false) {
		return errors.New("cannot move this piece to there")
	}

	ffirst := board.first[from.Y][from.X]
	tfirst := board.first[to.Y][to.X]

	isCastling := piece.King.IsSymbol(fsymbol) && (diff.X == -2 || diff.X == 2)
	canCastling := isCastling && board.CanCastling(from, to)
	if isCastling && canCastling == false {
		return errors.New("cannnot castle to that point")
	} else if canCastling == true {
		board.Castling(from, to)
	} else {
		board.Matrix[to.Y][to.X] = board.Matrix[from.Y][from.X]
		board.Matrix[from.Y][from.X] = ' '
		board.first[from.Y][from.X] = false
		board.first[to.Y][to.X] = false
	}
	if board.IsChecked(c) {
		board.Matrix[from.Y][from.X] = fsymbol
		board.Matrix[to.Y][to.X] = tsymbol
		board.first[from.Y][from.X] = ffirst
		board.first[to.Y][to.X] = tfirst
		return errors.New("cannot move that piece there: your king will be checked")
	}

	return nil
}
