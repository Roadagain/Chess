package point

type Point struct {
	Y int
	X int
}

func (point Point) Diff(to Point) Point {
	return Point{to.Y - point.Y, to.X - point.X}
}

func (point *Point) StepTo(to Point) {
	if point.Y < to.Y {
		point.Y++
	} else if point.Y > to.Y {
		point.Y--
	}
	if point.X < to.X {
		point.X++
	} else if point.X > to.X {
		point.X--
	}
}
