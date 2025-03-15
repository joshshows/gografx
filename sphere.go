package main

import (
	"image/color"
	"math"
)

type Sphere struct {
	center Vector
	radius float64
	color  color.RGBA
}

func (s Sphere) IntersectsAt(origin Vector, direction Vector) (bool, Vector) {
	L := origin.Subtract(s.center)
	b := 2.0 * direction.DotProduct(L)
	c := L.DotProduct(L) - s.radius*s.radius

	discriminant := b*b - 4*c

	if discriminant < 0 {
		// No intersection
		return false, Vector{}
	}

	// Compute the two intersection points
	sqrtD := math.Sqrt(float64(discriminant))
	t1 := (-b - sqrtD) / 2
	//t2 := (-b + sqrtD) / 2

	// Compute intersection points
	intersection1 := origin.Add(direction.Multiply(t1))
	//intersection2 := origin.Add(direction.Multiply(t2)) // need to figure out which is actually closer?

	return true, intersection1
}

func (s Sphere) Color() color.RGBA {
	return s.color
}

func (s Sphere) GetTexture(v Vector) color.RGBA {
	// Sin calc should give us a smooth change between 0 and 1
	factor := (math.Sin(v.x+v.y+v.z) + 1) / 2
	newBlue := uint8(0)
	newRed := uint8(float64(s.color.R) * factor)
	newGreen := uint8(float64(s.color.G) * factor)
	return color.RGBA{newRed, newGreen, newBlue, 255}
}

func (s Sphere) GetSurfaceNormal(surfacePoint Vector) Vector {
	v := surfacePoint.Subtract(s.center)
	v.Normalize()
	return v
}
