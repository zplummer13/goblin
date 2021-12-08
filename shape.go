package goblin

// Shape represents a convex polygon. Position gives the position relative to the space that it belongs to. The points
// are given in clockwise order and are relative to the position.
type Shape interface {
}

var _ Shape = &Polygon{}

type Polygon struct {
	Position Vector
	Points   []Vector
}

func NewRect(x float64, y float64, width float64, height float64) *Polygon {
	points := []Vector{
		{0, 0},
		{width, 0},
		{width, height},
		{0, height},
	}
	return &Polygon{
		Position: Vector{x, y},
		Points:   points,
	}
}
