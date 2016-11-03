package main

import (
	"board"
	"color"
	"enemy"
	"fmt"
	"io"
	"matrix"
	"os"
)

func scanMove() (*matrix.Move, error) {
	var file byte
	var rank int

	_, err := fmt.Scanf("%c%d", &file, &rank)
	return matrix.NewMove(file, rank), err
}

func main() {
	var player color.Color
	if len(os.Args) == 1 {
		player = color.White
	} else {
		player = color.ParseColor(os.Args[1])
		if player == color.Unknown {
			fmt.Printf("Invalid color: %s\n", os.Args[1])
			return
		}
	}
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
			if now != player {
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
			canMove, err := chessboard.CanMove(from, to, now)
			success = canMove
			if canMove {
				success = true
				chessboard.Move(from, to, chessboard.IsCastling(from, to))
			} else {
				fmt.Println(err)
			}
		}
		fmt.Println()
		now = now.Enemy()
	}
}
