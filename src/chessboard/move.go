package chessboard

type Move struct {
	file byte
	rank int
}

func NewMove(file byte, rank int) *Move {
	move := new(Move)
	move.file = file
	move.rank = rank
	return move
}

func (move Move) ToPoint() Point {
	return Point{int(move.file - 'a'), 8 - move.rank}
}
