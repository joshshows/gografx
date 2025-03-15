package main

import (
	"image/color"
	"math"
)

type Plane struct {
	point  Vector
	normal Vector
	color  color.RGBA
}

func (p Plane) IntersectsAt(origin Vector, direction Vector) (bool, Vector) {

	// https://en.wikipedia.org/wiki/Line%E2%80%93plane_intersection
	p.normal.Normalize() // Ensure the normal is unit length

	denominator := direction.DotProduct(p.normal)

	// Check if the ray is parallel to the plane
	// This is dealing with goofy rounding errors.
	if math.Abs(float64(denominator)) < 1e-6 {
		return false, Vector{}
	}

	numerator := p.point.Subtract(origin).DotProduct(p.normal)
	d := numerator / denominator

	// If d < 0, intersection is behind the ray's origin
	if d < 0 {
		return false, Vector{}
	}

	intersection := origin.Add(direction.Multiply(d))
	return true, intersection
}

func (p Plane) Color() color.RGBA {
	return p.color
}

func (p Plane) GetSurfaceNormal(surfacePoint Vector) Vector {
	v := p.normal.Subtract(surfacePoint)
	v.Normalize()
	return v
}
