package board

import (
	"color"
	"errors"
	"fmt"
	"matrix"
	"piece"
	"point"
)

type Board struct {
	matrix matrix.Matrix
	first  [8][8]bool
}

func NewBoard() *Board {
	board := new(Board)
	board.matrix = matrix.Starting()
	for i := 0; i < 8; i++ {
		board.first[0][i] = true
		board.first[1][i] = true
		board.first[6][i] = true
		board.first[7][i] = true
	}
	return board
}

func (board *Board) Print() {
	fmt.Print(" ")
	for i := 0; i < 8; i++ {
		fmt.Printf(" %c", 'a'+i)
	}
	fmt.Println()

	for i := 0; i < 8; i++ {
		fmt.Print(" +")
		for j := 0; j < 8; j++ {
			fmt.Print("-+")
		}
		fmt.Println()
		fmt.Printf("%d|", 8-i)
		for j := 0; j < 8; j++ {
			fmt.Printf("%c|", board.matrix[i][j])
		}
		fmt.Println(8 - i)
	}
	fmt.Print(" +")
	for i := 0; i < 8; i++ {
		fmt.Print("-+")
	}
	fmt.Println()

	fmt.Print(" ")
	for i := 0; i < 8; i++ {
		fmt.Printf(" %c", 'a'+i)
	}
	fmt.Println()
}

func (board Board) IsChecked(c color.Color) bool {
	var king point.Point
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			symbol := board.matrix[i][j]
			if piece.King.IsSymbol(symbol) && color.WhichColor(symbol) == c {
				king = point.Point{i, j}
				break
			}
		}
	}
	enemy := c.Enemy()
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			symbol := board.matrix[i][j]
			if color.WhichColor(symbol) != enemy {
				continue
			}

			from := point.Point{i, j}
			piece := piece.WhichPiece(symbol)
			first := board.first[i][j]
			diff := from.Diff(king)
			if piece.CanMove(diff, first, true) && board.matrix.ExistBarrier(from, king) == false {
				return true
			}
		}
	}
	return false
}

func (board *Board) CanCastling(from, to point.Point) bool {
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
	fsymbol := board.matrix[rfrom.Y][rfrom.X]
	if piece.Rook.IsSymbol(fsymbol) == false {
		return false
	} else if board.first[rfrom.Y][rfrom.X] == false {
		return false
	} else if board.matrix.ExistBarrier(rfrom, rto) {
		return false
	}
	return true
}

func (board *Board) Castling(from, to point.Point) {
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

	board.matrix[to.Y][to.X] = board.matrix[from.Y][from.X]
	board.matrix[from.Y][from.X] = ' '
	board.first[from.Y][from.X] = false
	board.first[to.Y][to.X] = false
	board.matrix[rto.Y][rto.X] = board.matrix[rfrom.Y][rfrom.X]
	board.matrix[rfrom.Y][rfrom.X] = ' '
	board.first[rfrom.Y][rfrom.X] = false
	board.first[rto.Y][rto.X] = false
}

func (board *Board) Move(from, to point.Point, c color.Color) error {
	if matrix.InMatrix(from) == false || matrix.InMatrix(to) == false {
		return errors.New("out of board")
	}

	fsymbol := board.matrix[from.Y][from.X]
	tsymbol := board.matrix[to.Y][to.X]

	fcolor := color.WhichColor(board.matrix[from.Y][from.X])
	tcolor := color.WhichColor(board.matrix[to.Y][to.X])
	diff := point.Point{0, 0}
	if fcolor == color.White {
		diff = from.Diff(to)
	} else if fcolor == color.Black {
		diff = to.Diff(from)
	}
	if fcolor != c || tcolor == c {
		return errors.New("cannot move this piece")
	}

	fpiece := piece.WhichPiece(fsymbol)
	toEnemy := color.WhichColor(board.matrix[to.Y][to.X]) == c.Enemy()
	canMove := fpiece.CanMove(diff, board.first[from.Y][from.X], toEnemy)
	existBarrier := board.matrix.ExistBarrier(from, to)
	if canMove == false || (existBarrier && piece.Knight.IsSymbol(fsymbol) == false) {
		return errors.New("cannot move this piece to that point")
	}

	ffirst := board.first[from.Y][from.X]
	tfirst := board.first[to.Y][to.X]

	isCastling := piece.King.IsSymbol(fsymbol) && (diff.X == -2 || diff.X == 2)
	canCastling := isCastling && board.CanCastling(from, to)
	if isCastling && canCastling == false {
		return errors.New("cannot move this piece to that point")
	} else if canCastling == true {
		board.Castling(from, to)
	} else {
		board.matrix[to.Y][to.X] = board.matrix[from.Y][from.X]
		board.matrix[from.Y][from.X] = ' '
		board.first[from.Y][from.X] = false
		board.first[to.Y][to.X] = false
	}
	if board.IsChecked(c) {
		board.matrix[from.Y][from.X] = fsymbol
		board.matrix[to.Y][to.X] = tsymbol
		board.first[from.Y][from.X] = ffirst
		board.first[to.Y][to.X] = tfirst
		return errors.New("cannot move this piece to that point")
	}

	return nil
}
