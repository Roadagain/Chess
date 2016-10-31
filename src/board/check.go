package board

import (
	"color"
	"matrix"
	"piece"
)

func (board Board) IsChecked(c color.Color) bool {
	var king matrix.Point
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			symbol := board.matrix[i][j]
			if piece.King.IsSymbol(symbol) && color.WhichColor(symbol) == c {
				king = matrix.Point{i, j}
				break
			}
		}
	}
	enemy := c.Enemy()
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			symbol := board.matrix[i][j]
			if color.WhichColor(symbol) != enemy {
				continue
			}

			from := matrix.Point{i, j}
			piece := piece.WhichPiece(symbol)
			first := board.first[i][j]
			diff := from.Diff(king)
			if piece.CanMove(diff, first, true) && board.matrix.ExistBarrier(from, king) == false {
				return true
			}
		}
	}
	return false
}

func (board Board) IsCheckMate(c color.Color) bool {
	if board.IsChecked(c) == false {
		return false
	}

	friend := make([]matrix.Point, 16)
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			symbol := board.matrix[i][j]
			if color.WhichColor(symbol) == c {
				friend = append(friend, matrix.Point{i, j})
			}
		}
	}

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			to := matrix.Point{i, j}
			for _, k := range friend {
				fsymbol := board.matrix[k.Y][k.X]
				tsymbol := board.matrix[to.Y][to.X]
				tcolor := color.WhichColor(tsymbol)
				if tcolor == c || tcolor == color.Unknown {
					continue
				}

				piece := piece.WhichPiece(fsymbol)
				first := board.first[k.Y][k.X]
				diff := k.Diff(to)
				toEnemy := color.WhichColor(tsymbol) == c.Enemy()
				if piece.CanMove(diff, first, toEnemy) && board.matrix.ExistBarrier(k, to) == false {
					board.matrix[to.Y][to.X] = fsymbol
					board.matrix[k.Y][k.X] = ' '
					if board.IsChecked(c) == false {
						board.Print()
						board.matrix[to.Y][to.X] = tsymbol
						board.matrix[k.Y][k.X] = fsymbol
						return false
					}
					board.matrix[to.Y][to.X] = tsymbol
					board.matrix[k.Y][k.X] = fsymbol
				}
			}
		}
	}
	return true
}
