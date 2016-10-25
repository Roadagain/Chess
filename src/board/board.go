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
}

func NewBoard() *Board {
	board := new(Board)
	board.matrix = matrix.Starting()
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

func (board *Board) Move(from, to point.Point, c color.Color) error {
	if matrix.InMatrix(from) == false || matrix.InMatrix(to) == false {
		return errors.New("out of board")
	}

	fsymbol := board.matrix[from.Y][from.X]
	// tsymbol := board.matrix[to.Y][to.X]

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
	canMove := fpiece.CanMove(diff)
	existBarrier := board.matrix.ExistBarrier(from, to)
	if canMove == false || (existBarrier && piece.Knight.IsSymbol(fsymbol) == false) {
		return errors.New("cannot move this piece to that point")
	}

	board.matrix[to.Y][to.X] = board.matrix[from.Y][from.X]
	board.matrix[from.Y][from.X] = ' '

	return nil
}
