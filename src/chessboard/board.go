package chessboard

import (
	"errors"
	"fmt"
)

type Color int

const (
	Unknown Color = iota
	Empty
	White
	Black
)

func (color Color) String() string {
	switch color {
	case Empty:
		return "Empty"
	case White:
		return "White"
	case Black:
		return "Black"
	default:
		return "Unknown"
	}
}

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

func (board Board) in(point Point) bool {
	return 0 <= point.x && point.x < 8 && 0 <= point.y && point.y < 8
}

func (board Board) color(point Point) (Color, error) {
	if board.in(point) == false {
		return Unknown, errors.New("out of board")
	}

	piece := board[point.y][point.x]
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

func (board *Board) Move(from, to Point) error {
	if board.in(from) == false || board.in(to) == false {
		return errors.New("out of board")
	} else if board.matrix[from.y][from.x] == ' ' {
		return errors.New("cannot move empty piece")
	}
	board.matrix[to.y][to.x] = board.matrix[from.y][from.x]
	board.matrix[from.y][from.x] = ' '

	return nil
}
