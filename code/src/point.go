package main

import (
	"animal"
	"fmt"
	"point"
)

func main() {

	//tag::NewPoint[]
	p := *new(point.Point)
	p.X = 1
	p.Y = 2
	//end::NewPoint[]

	//tag::PointLiteral[]
	p = point.Point{X: 1, Y: 2}
	//end::PointLiteral[]

	//tag::TranslatingPoints[]
	q := p.Translate(5, 5)
	fmt.Printf("Translated %v to %v\n", p, q)
	//end::TranslatingPoints[]

	//tag::TranslatingPointers[]
	qP := &point.Point{X: 1, Y: 2}
	qP.TranslatePointer(5, 5)
	fmt.Printf("Translated using pointer to %v\n", qP)
	//end::TranslatingPointers[]

	//tag::PrintingPointAndColorPoint[]
	r := point.ColorPoint{Point: point.Point{X: 1, Y: 4}, Color: point.BLUE}

	fmt.Printf("Point: %v\n", p)
	fmt.Printf("Color Point: %v\n", r)
	//end::PrintingPointAndColorPoint[]

	//tag::CalculatingDistance[]
	fmt.Printf("Dist b/w p and q = %v\n", p.DistanceTo(r))
	fmt.Printf("Dist b/w q and p = %v\n", r.DistanceTo(p))
	//end::CalculatingDistance[]

	//tag::CalculatingDistanceWithAnimals[]
	penguin := animal.Animal{Name: "penguin", X: 1, Y: 1}
	seal := animal.Animal{Name: "seal", X: 1, Y: 4}

	fmt.Printf("Dist b/w penguin and seal = %v\n", penguin.DistanceTo(seal))
	fmt.Printf("Dist b/w penguin and point = %v\n", penguin.DistanceTo(p))
	//end::CalculatingDistanceWithAnimals[]
}
