package board

import "fmt"

func (board Board) Print() {
	fmt.Println(board.Matrix.ToString())
}
