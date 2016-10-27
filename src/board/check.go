package board

import (
	"color"
	"piece"
	"point"
)

func (board Board) IsChecked(c color.Color) bool {
	var king point.Point
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			symbol := board.matrix[i][j]
			if piece.King.IsSymbol(symbol) && color.WhichColor(symbol) == c {
				king = point.Point{i, j}
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

			from := point.Point{i, j}
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
