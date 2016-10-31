package board

import (
	"fmt"
	"matrix"
)

func (board *Board) Print() {
	fmt.Print(" ")
	for i := 0; i < matrix.SIDE; i++ {
		fmt.Printf(" %c", 'a'+i)
	}
	fmt.Println()

	for i := 0; i < matrix.SIDE; i++ {
		fmt.Print(" +")
		for j := 0; j < matrix.SIDE; j++ {
			fmt.Print("-+")
		}
		fmt.Println()
		fmt.Printf("%d|", matrix.SIDE-i)
		for j := 0; j < matrix.SIDE; j++ {
			fmt.Printf("%c|", board.matrix[i][j])
		}
		fmt.Println(matrix.SIDE - i)
	}
	fmt.Print(" +")
	for i := 0; i < matrix.SIDE; i++ {
		fmt.Print("-+")
	}
	fmt.Println()

	fmt.Print(" ")
	for i := 0; i < matrix.SIDE; i++ {
		fmt.Printf(" %c", 'a'+i)
	}
	fmt.Println()
}
