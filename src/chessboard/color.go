package chessboard

type Color int

const (
	Unknown Color = iota
	Empty
	White
	Black
)

func (color Color) String() string {
	switch color {
	case Empty:
		return "Empty"
	case White:
		return "White"
	case Black:
		return "Black"
	default:
		return "Unknown"
	}
}

func (color Color) Enemy() Color {
	switch color {
	case White:
		return Black
	case Black:
		return White
	default:
		return color
	}
}
