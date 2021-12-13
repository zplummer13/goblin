package goblin

import (
	"math"
)

type Actor struct {
	Position Vector
	Box      *Box
	Space    *Space
}

func NewActor(x float64, y float64, width float64, height float64) *Actor {
	return &Actor{
		Position: Vector{x, y},
		Box:      NewBox(x, y, width, height),
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

	updatedBox := *a.Box
	updatedBox.Position = newPosition

	solids := []*Solid{}
	for _, solid := range a.Space.Solids {
		if solid.Box.CheckCollision(&updatedBox) {
			solids = append(solids, solid)
		}
	}

	actors := []*Actor{}
	for _, actor := range a.Space.Actors {
		if actor.Box.CheckCollision(&updatedBox) {
			actors = append(actors, actor)
		}
	}

	if len(solids) > 0 || len(actors) > 0 {
		return &Collision{
			CheckActor: a,
			Position:   newPosition,
			Actors:     actors,
			Solids:     solids,
		}
	}

	return nil
}
