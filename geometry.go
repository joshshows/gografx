package main

import "image/color"

type geometry interface {
	IntersectsAt(origin Vector, direction Vector) (bool, Vector)
	Color(point Vector) color.RGBA
	GetSurfaceNormal(surfacePoint Vector) Vector
}

func MultiplyColor(c color.RGBA, x float64) color.RGBA {
	fac := [3]float64{float64(c.R), float64(c.G), float64(c.B)}
	return color.RGBA{uint8(fac[0] * x), uint8(fac[1] * x), uint8(fac[2] * x), 255}
}
