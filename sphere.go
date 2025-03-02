package main

import (
	"math"
)

type Sphere struct {
	center Vector
	radius float64
}

func (s Sphere) IntersectsAt(origin Vector, direction Vector) (bool, Vector, Vector) {
	L := origin.Subtract(s.center)
	b := 2.0 * direction.DotProduct(L)
	c := L.DotProduct(L) - s.radius*s.radius

	discriminant := b*b - 4*c

	if discriminant < 0 {
		// No intersection
		return false, Vector{}, Vector{}
	}

	// Compute the two intersection points
	sqrtD := math.Sqrt(float64(discriminant))
	t1 := (-b - sqrtD) / 2
	t2 := (-b + sqrtD) / 2

	// Compute intersection points
	intersection1 := origin.Add(direction.Multiply(t1))
	intersection2 := origin.Add(direction.Multiply(t2))

	return true, intersection1, intersection2
}
