package point

import "math"

//tag::PointStruct[]
type Point struct {
	X, Y float64
}

//end::PointStruct[]

//tag::ColorPoint[]
const (
	BLUE  = iota
	RED   = iota
	GREEN = iota
)

type ColorPoint struct {
	Point Point
	Color int
}

//end::ColorPoint[]

//tag::Translate[]
func (p Point) Translate(xDist float64, yDist float64) Point {
	return Point{p.X + xDist, p.Y + yDist}
}

//end::Translate[]

//tag::TranslatePointer[]
func (p *Point) TranslatePointer(xDist float64, yDist float64) {
	p.X = p.X + xDist
	p.Y = p.Y + yDist
}

//end::TranslatePointer[]

//tag::Interfaces[]
type Positioner interface {
	Coordinates() Point
}

type Distancer interface {
	DistanceTo(p Positioner) float64
}

//end::Interfaces[]

//tag::PointSatisfiesInt[]
func (p Point) Coordinates() Point {
	return p
}

func (p Point) DistanceTo(pos Positioner) float64 {
	return distanceBetween(p, pos)
}

//end::PointSatisfiesInt[]

//tag::ColorPointSatisfiesInt[]
func (cp ColorPoint) Coordinates() Point {
	return cp.Point
}

func (cp ColorPoint) DistanceTo(pos Positioner) float64 {
	return distanceBetween(cp, pos)
}

//end::ColorPointSatisfiesInt[]

//tag::distanceBetween[]
func distanceBetween(a Positioner, b Positioner) float64 {
	p := a.Coordinates()
	q := b.Coordinates()
	sqOfXDist := math.Pow(p.X-q.X, 2)
	sqOfYDist := math.Pow(p.Y-q.Y, 2)
	return math.Sqrt(sqOfXDist + sqOfYDist)
}

//end::distanceBetween[]
