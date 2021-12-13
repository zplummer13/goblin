package goblin

type Solid struct {
	Space    *Space
	Position Vector
	Box      *Box
}

func NewSolid() *Solid {
	return &Solid{}
}
