package main

import (
	"board"
	"color"
	"fmt"
	"point"
)

func scanMove() (*point.Move, error) {
	var file byte
	var rank int

	_, err := fmt.Scanf("%c%d", &file, &rank)
	return point.NewMove(file, rank), err
}

func main() {
	chessboard := board.NewBoard()
	finish := false
	now := color.White

	for finish == false {
		success := false

		chessboard.Print()
		if chessboard.IsCheckMate(now) {
			fmt.Println("Checkmate")
			break
		} else if chessboard.IsChecked(now) {
			fmt.Println("Your king is checked")
		}

		for success == false {
			from, err := scanMove()
			if err != nil {
				finish = true
				break
			}
			to, err := scanMove()
			if err != nil {
				finish = true
				break
			}
			err = chessboard.Move(from.ToPoint(), to.ToPoint(), now)
			success = err == nil
			if success == false {
				fmt.Println(err)
			}
		}
		fmt.Println()
		now = now.Enemy()
	}
}
