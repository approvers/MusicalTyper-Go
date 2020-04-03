package pos

// Pos expresses position on screen
type Pos struct {
	x int
	y int
}

// FromXY makes Pos from x and y coordinates
func FromXY(x, y int) Pos {
	return Pos{x: x, y: y}
}

// X returns x coordinate of Pos
func (p Pos) X() int {
	return p.x
}

// Y returns y coordinate of Pos
func (p Pos) Y() int {
	return p.y
}
