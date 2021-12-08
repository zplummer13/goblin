package goblin

import "math"

type Actor struct {
	Position Vector
	Shape    Shape
	Space    *Space
}

func NewActor(x float64, y float64, shape Shape) *Actor {
	return &Actor{
		Position: Vector{x, y},
		Shape:    shape,
	}
}

func (a *Actor) SetSpace(s *Space) {
	a.Space = s
}

func (a *Actor) MoveX(amount float64, onCollide func(c *Collision)) {

	move := int(math.Round(amount))

	if move != 0 {

		sign := Vector{1, 0}
		if math.Signbit(float64(move)) {
			sign = Vector{-1, 0}
		}

		for move != 0 {
			collision := a.collidesAt(a.Position.Add(sign))
			if collision == nil {
				// There is no Solid immediately beside us
				a.Position = a.Position.Add(sign)
				move -= int(sign.X())
			} else {
				// Hit a solid!
				if onCollide != nil {
					onCollide(collision)
				}
				break
			}
		}
	}
}

func (a *Actor) MoveY(amount float64, onCollide func(c *Collision)) {

	move := int(math.Round(amount))

	if move != 0 {

		sign := Vector{0, 1}
		if math.Signbit(float64(move)) {
			sign = Vector{0, -1}
		}

		for move != 0 {
			collision := a.collidesAt(a.Position.Add(sign))
			if collision == nil {
				// There is no Solid immediately beside us
				a.Position = a.Position.Add(sign)
				move -= int(sign.Y())
			} else {
				// Hit a solid!
				if onCollide != nil {
					onCollide(collision)
				}
				break
			}
		}
	}
}

func (a *Actor) collidesAt(newPosition Vector) *Collision {

	// for _, solid := range a.Space.Solids {

	// }

	// for _, actor := range a.Space.Actors {

	// }

	return nil
}
