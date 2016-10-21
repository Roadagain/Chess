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
	finish := false

	for finish == false {
		var from *chessboard.Move
		var to *chessboard.Move
		var err error
		success := false

		board.Print()
		for success == false {
			from, err = scanMove()
			if err != nil {
				finish = true
				break
			}
			to, err = scanMove()
			if err != nil {
				finish = true
				break
			}
			err = board.Move(from.ToPoint(), to.ToPoint())
			success = err == nil
			if success == false {
				fmt.Println(err)
			}
		}
		fmt.Println()
	}
}
