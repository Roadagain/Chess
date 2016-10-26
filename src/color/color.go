package color

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

func WhichColor(c byte) Color {
	if 'A' <= c && c <= 'Z' {
		return White
	} else if 'a' <= c && c <= 'z' {
		return Black
	} else if c == ' ' {
		return Empty
	} else {
		return Unknown
	}
}
