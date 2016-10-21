package chessboard

var starting = [][]byte{
	[]byte{'R', 'N', 'B', 'K', 'Q', 'B', 'N', 'R'},
	[]byte{'P', 'P', 'P', 'P', 'P', 'P', 'P', 'P'},
	[]byte{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
	[]byte{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
	[]byte{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
	[]byte{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
	[]byte{'P', 'P', 'P', 'P', 'P', 'P', 'P', 'P'},
	[]byte{'R', 'N', 'B', 'K', 'Q', 'B', 'N', 'R'},
}

func Piece(p Point) byte {
	return starting[p.y][p.x]
}
