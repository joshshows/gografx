package main

import (
	"image/color"
	"math"
)

type CheckerPlane struct {
	Plane
}

func (p CheckerPlane) Color(point Vector) color.RGBA {
	//return color.RGBA{0, 0, 255, 255}
	// Find two perpendicular vectors (U, V) in the plane
	uAxis, vAxis := p.ComputeUVAxes()

	// Project intersection point onto these axes
	u := int(math.Floor(float64(point.DotProduct(uAxis) * 0.0015))) // Scale for pattern size
	v := int(math.Floor(float64(point.DotProduct(vAxis) * 0.0015)))

	// Check if it's an even or odd square
	if (u+v)%2 == 0 {
		return color.RGBA{255, 255, 255, 0}
	} else {
		return color.RGBA{0, 0, 0, 0}
	}
}
