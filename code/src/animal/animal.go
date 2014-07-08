package animal

import (
	"math"
	"point"
)

//tag::Animal[]
type Animal struct {
	Name string
	X, Y float64
}

//end::Animal[]

//tag::AnimalSatisfiesDP[]
func (a Animal) Coordinates() point.Point {
	return point.Point{X: a.X, Y: a.Y}
}

func (a Animal) DistanceTo(pos point.Positioner) float64 {
	thing := pos.Coordinates()
	sqOfXDist := math.Pow(a.X-thing.X, 2)
	sqOfYDist := math.Pow(a.Y-thing.Y, 2)
	return math.Sqrt(sqOfXDist + sqOfYDist)
}

//end::AnimalSatisfiesDP[]
