package goblin

import "math"

const (
	x = iota
	y
	z
)

type Vector []float64

func (v Vector) X() float64 {
	if len(v) > 0 {
		return v[x]
	}
	return 0
}

func (v Vector) Y() float64 {
	if len(v) > 1 {
		return v[y]
	}
	return 0
}

func (v Vector) Add(v2 Vector) Vector {
	return Vector{v.X() + v2.X(), v.Y() + v2.Y()}
}

func (v Vector) Sub(v2 Vector) Vector {
	return Vector{v.X() - v2.X(), v.Y() - v2.Y()}
}

func (v Vector) Equal(v2 Vector) bool {
	if len(v) != len(v2) {
		return false
	}
	return v.X() == v2.X() && v.Y() == v2.Y()
}

func (v Vector) Magnitude() float64 {
	return math.Sqrt(v.X()*v.X() + v.Y()*v.Y())
}

// Use this instead of Magnitude when possible for performance reasons
func (v Vector) MagnitudeSquared() float64 {
	return v.X()*v.X() + v.Y()*v.Y()
}

func (v Vector) Scale(s float64) Vector {
	return Vector{v.X() * s, v.Y() * s}
}

func (v Vector) Dot(v2 Vector) float64 {
	return v.X()*v2.X() + v.Y()*v2.Y()
}
