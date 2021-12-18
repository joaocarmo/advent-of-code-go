package helpers

type Point struct {
	X int
	Y int
	Z int
}

func (p *Point) New(x int, y int, z int) {
	p.X = x
	p.Y = y
	p.Z = z
}
