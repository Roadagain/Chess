package main

import (
	"chessboard"
	"fmt"
)

func scanMove() (*chessboard.Move, error) {
	var file byte
	var rank int

	_, err := fmt.Scanf("%c%d", &file, &rank)
	return chessboard.NewMove(file, rank), err
}

func main() {
	board := chessboard.NewBoard()

	for true {
		var from *chessboard.Move
		var to *chessboard.Move
		var err error

		board.Print()
		from, err = scanMove()
		if err != nil {
			break
		}
		to, err = scanMove()
		if err != nil {
			break
		}
		err = board.Move(from.ToPoint(), to.ToPoint())
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println()
	}
}
