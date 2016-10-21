package main

import (
	"chessboard"
	"fmt"
)

func scanMove() (*chessboard.Move, error) {
	var file byte
	var rank int

	_, err := fmt.Scanf("%c %d", &file, &rank)
	return chessboard.NewMove(file, rank), err
}

func main() {
	board := chessboard.NewBoard()

	for true {
		board.Print()
		from, err := scanMove()
		if err != nil {
			break
		}
		to, _ := scanMove()
		board.Move(from.ToPoint(), to.ToPoint())
	}
}
