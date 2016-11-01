package board

import (
	"fmt"
	"matrix"
)

func (board Board) Print() {
	fmt.Println(board.matrix.ToString())
}
