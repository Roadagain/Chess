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
			symbol := board.Matrix[i][j]
			if piece.King.IsSymbol(symbol) && color.WhichColor(symbol) == c {
				king = matrix.Point{i, j}
				break
			}
		}
	}
	enemy := c.Enemy()
	enemies := board.Matrix.Positions(enemy)
	for _, i := range enemies {
		symbol := board.Matrix[i.Y][i.X]
		p := piece.WhichPiece(symbol)
		first := board.first[i.Y][i.X]
		var diff matrix.Point
		if c == color.White {
			diff = i.Diff(king)
		} else {
			diff = king.Diff(i)
		}

		canMove := p.CanMove(diff, first, true)
		existBarrier := board.Matrix.ExistBarrier(i, king)
		if canMove && (piece.Knight.IsSymbol(symbol) || existBarrier == false) {
			return true
		}
	}
	return false
}

func (board Board) IsCheckMate(c color.Color) bool {
	if board.IsChecked(c) == false {
		return false
	}

	friend := board.Matrix.Positions(c)
	movable := append(board.Matrix.Positions(color.Empty), board.Matrix.Positions(c.Enemy())...)

	for _, i := range movable {
		for _, j := range friend {
			fsymbol := board.Matrix[j.Y][j.X]
			tsymbol := board.Matrix[i.Y][i.X]
			tcolor := color.WhichColor(tsymbol)
			if tcolor == c || tcolor == color.Unknown {
				continue
			}

			piece := piece.WhichPiece(fsymbol)
			first := board.first[j.Y][j.X]
			var diff matrix.Point
			if c == color.White {
				diff = j.Diff(i)
			} else if c == color.Black {
				diff = i.Diff(j)
			}
			toEnemy := color.WhichColor(tsymbol) == c.Enemy()
			if piece.CanMove(diff, first, toEnemy) && board.Matrix.ExistBarrier(j, i) == false {
				board.Matrix[i.Y][i.X] = fsymbol
				board.Matrix[j.Y][j.X] = ' '
				if board.IsChecked(c) == false {
					board.Matrix[i.Y][i.X] = tsymbol
					board.Matrix[j.Y][j.X] = fsymbol
					return false
				}
				board.Matrix[i.Y][i.X] = tsymbol
				board.Matrix[j.Y][j.X] = fsymbol
			}
		}
	}
	return true
}
