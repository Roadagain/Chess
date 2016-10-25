package main

import (
	"chessboard"
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
	board := chessboard.NewBoard()
	finish := false
	now := color.White

	for finish == false {
		success := false

		board.Print()
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
			err = board.Move(from.ToPoint(), to.ToPoint(), now)
			success = err == nil
			if success == false {
				fmt.Println(err)
			}
		}
		fmt.Println()
		now = now.Enemy()
	}
}
