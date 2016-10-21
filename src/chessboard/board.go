package chessboard

import "fmt"

var starting = [8][8]byte{
	{'R', 'N', 'B', 'K', 'Q', 'B', 'N', 'R'},
	{'P', 'P', 'P', 'P', 'P', 'P', 'P', 'P'},
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

func (board *Board) Move(from, to Point) {
	board.matrix[to.y][to.x] = board.matrix[from.y][from.x]
	board.matrix[from.y][from.x] = ' '
}
