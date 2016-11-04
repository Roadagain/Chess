package main

import (
	"board"
	"color"
	"enemy"
	"flag"
	"fmt"
	"io"
	"matrix"
	"piece"
)

func scanMove() (*matrix.Move, error) {
	var file byte
	var rank int

	_, err := fmt.Scanf("%c%d", &file, &rank)
	return matrix.NewMove(file, rank), err
}

func main() {
	enemyFlag := flag.String("enemy", "random", "Enemy")
	flag.Parse()

	enemyType := enemy.ParseType(*enemyFlag)
	if enemyType == enemy.Unknown {
		fmt.Printf("Invalid enemy: %s\n", *enemyFlag)
		return
	}

	var player color.Color
	args := flag.Args()
	if len(args) == 0 {
		player = color.White
	} else {
		player = color.ParseColor(args[1])
		if player == color.Unknown {
			fmt.Printf("Invalid color: %s\n", args[1])
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
				ene := enemy.NewEnemy(chessboard, now)
				switch enemyType {
				case enemy.Brutal:
					from, to = ene.BrutalSelect()
				default:
					from, to = ene.RandomizedSelect()
				}
				fmt.Printf("%c%d %c%d\n", byte(from.X+'a'), 8-from.Y, byte(to.X+'a'), 8-to.Y)
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
				isPawn := piece.Pawn.IsSymbol(chessboard.Matrix[to.Y][to.X])
				if isPawn && matrix.IsDeepest(to, now) {
					fmt.Print("Promotion: ")
					var choice string
					if now == player {
						fmt.Scan(&choice)
					} else {
						choice = "Queen"
						fmt.Println(choice)
					}

					p := piece.WhichPiece([]byte(choice)[0])
					for piece.CanPromotionTo(p) == false {
						fmt.Println("Invalid symbol")
						fmt.Print("Promotion: ")
						fmt.Scan(&choice)
						p = piece.WhichPiece([]byte(choice)[0])
					}
					chessboard.Promotion(to, p, now)
				}
			} else {
				fmt.Println(err)
			}
		}
		fmt.Println()
		now = now.Enemy()
	}
}
