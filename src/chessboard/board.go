package chessboard

var starting = [8][8]byte{
	{'R', 'N', 'B', 'K', 'Q', 'B', 'N', 'R'},
	{'P', 'P', 'P', 'P', 'P', 'P', 'P', 'P'},
	{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
	{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
	{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
	{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
	{'P', 'P', 'P', 'P', 'P', 'P', 'P', 'P'},
	{'R', 'N', 'B', 'K', 'Q', 'B', 'N', 'R'},
}

type Board struct {
	matrix [8][8]byte
}

func NewBoard() *Board {
	board := new(Board)
	board.matrix = starting
	return board
}

func (board Board) Piece(p Point) byte {
	return board.matrix[p.y][p.x]
}
