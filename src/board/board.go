package board

import (
	"color"
	"errors"
	"fmt"
	"piece"
	"point"
)

var starting = [8][8]byte{
	{'r', 'n', 'b', 'k', 'q', 'b', 'n', 'r'},
	{'p', 'p', 'p', 'p', 'p', 'p', 'p', 'p'},
	{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
	{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
	{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
	{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
	{'P', 'P', 'P', 'P', 'P', 'P', 'P', 'P'},
	{'R', 'N', 'B', 'K', 'Q', 'B', 'N', 'R'},
}

type Board struct {
	matrix [8][8]byte
}

func NewBoard() *Board {
	board := new(Board)
	board.matrix = starting
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

func InBoard(point point.Point) bool {
	return 0 <= point.X && point.X < 8 && 0 <= point.Y && point.Y < 8
}

func (board *Board) Move(from, to point.Point, c color.Color) error {
	if InBoard(from) == false || InBoard(to) == false {
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
	fpiece := piece.WhichPiece(fsymbol)
	canMove := fpiece.CanMove(diff)
	if fcolor != c || tcolor == c || canMove == false {
		return errors.New("cannot move this piece")
	}

	board.matrix[to.Y][to.X] = board.matrix[from.Y][from.X]
	board.matrix[from.Y][from.X] = ' '

	return nil
}
