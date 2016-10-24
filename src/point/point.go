package point

type Point struct {
	Y int
	X int
}

func (point Point) Diff(to Point) Point {
	return Point{to.Y - point.Y, to.X - point.X}
}
