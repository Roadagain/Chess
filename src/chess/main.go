package main

import (
	"board"
	"color"
	"enemy"
	"fmt"
	"io"
	"matrix"
)

func scanMove() (*matrix.Move, error) {
	var file byte
	var rank int

	_, err := fmt.Scanf("%c%d", &file, &rank)
	return matrix.NewMove(file, rank), err
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

		var from, to matrix.Point
		var err error
		for success == false {
			if now == color.Black {
				from, to = enemy.NewEnemy(chessboard, now).RandomizedSelect()
			} else {
				var move *matrix.Move
				move, err = scanMove()
				if err == io.EOF {
					finish = true
					break
				} else if err != nil {
					fmt.Println(err)
					continue
				} else {
					from = move.ToPoint()
				}
				move, err = scanMove()
				if err == io.EOF {
					finish = true
					break
				} else if err != nil {
					fmt.Println(err)
					continue
				} else {
					to = move.ToPoint()
				}
			}
			err = chessboard.Move(from, to, now)
			success = err == nil
			if success == false {
				fmt.Println(err)
			}
		}
		fmt.Println()
		now = now.Enemy()
	}
}
