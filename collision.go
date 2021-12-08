package goblin

type Collision struct {
	CheckActor *Actor
	Position   Vector
	Actors     []*Actor
	Solids     []*Solid
}
