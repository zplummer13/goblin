package goblin

type Solid struct {
	Space    *Space
	Position Vector
	Shape    Shape
}

func NewSolid() *Solid {
	return &Solid{}
}
