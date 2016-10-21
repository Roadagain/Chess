package chessboard

func Piece(p Point) byte {
	starting := [][]byte{
		[]byte{'R', 'N', 'B', 'K', 'Q', 'B', 'N', 'R'},
		[]byte{'P', 'P', 'P', 'P', 'P', 'P', 'P', 'P'},
		[]byte{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
		[]byte{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
		[]byte{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
		[]byte{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
		[]byte{'P', 'P', 'P', 'P', 'P', 'P', 'P', 'P'},
		[]byte{'R', 'N', 'B', 'K', 'Q', 'B', 'N', 'R'},
	}

	return starting[p.y][p.x]
}
