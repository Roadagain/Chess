package chessboard

import (
	"errors"
	"fmt"
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

func (board Board) color(point point.Point) (Color, error) {
	if InBoard(point) == false {
		return Unknown, errors.New("out of board")
	}

	piece := board.matrix[point.Y][point.X]
	if piece == ' ' {
		return Empty, nil
	} else if 'A' <= piece && piece <= 'Z' {
		return White, nil
	} else if 'a' <= piece && piece <= 'z' {
		return Black, nil
	} else {
		return Unknown, errors.New("Unknown color")
	}
}

func (board *Board) Move(from, to point.Point, color Color) error {
	if InBoard(from) == false || InBoard(to) == false {
		return errors.New("out of board")
	}

	fcolor, ferr := board.color(from)
	tcolor, terr := board.color(to)
	if fcolor != color || ferr != nil || tcolor == color || terr != nil {
		return errors.New("cannot move this piece")
	}

	board.matrix[to.Y][to.X] = board.matrix[from.Y][from.X]
	board.matrix[from.Y][from.X] = ' '

	return nil
}
