package goblin

type Box struct {
	Position Vector
	Width    float64
	Height   float64
}

func NewBox(x float64, y float64, width float64, height float64) *Box {
	return &Box{
		Position: Vector{x, y},
		Width:    width,
		Height:   height,
	}
}

func (b *Box) X() float64 {
	return b.Position.X()
}

func (b *Box) Y() float64 {
	return b.Position.Y()
}

func (b *Box) CheckCollision(a *Box) bool {

	return a.X() < b.X()+b.Width &&
		a.X()+a.Width > b.X() &&
		a.Y() < b.Y()+b.Height &&
		a.Y()+a.Height > b.Y()
}
