package main

import "image/color"

type geometry interface {
	IntersectsAt(origin Vector, direction Vector) (bool, Vector)
	Color() color.RGBA
	GetSurfaceNormal(surfacePoint Vector) Vector
}
