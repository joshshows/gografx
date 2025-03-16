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

func (p Plane) Color(point Vector) color.RGBA {
	//return color.RGBA{0, 0, 255, 255}
	// Find two perpendicular vectors (U, V) in the plane
	uAxis, vAxis := p.ComputeUVAxes()

	// Project intersection point onto these axes
	/* checkerboard
	u := int(math.Floor(float64(point.DotProduct(uAxis) * 0.0015))) // Scale for pattern size
	v := int(math.Floor(float64(point.DotProduct(vAxis) * 0.0005)))

	// Check if it's an even or odd square
	if (u+v)%2 == 0 {
		return color.RGBA{255, 255, 255, 0}
	} else {
		return color.RGBA{0, 0, 0, 0}
	}
	*/

	frequency := 0.002
	u := point.DotProduct(uAxis)
	v := point.DotProduct(vAxis)

	intensity := 0.5 * (1 + float32(math.Sin(float64(frequency*u))*math.Sin(float64(frequency*v))))

	return MultiplyColor(p.color, float64(intensity))
}

func (p Plane) GetSurfaceNormal(surfacePoint Vector) Vector {
	v := p.normal.Subtract(surfacePoint)
	v.Normalize()
	return v
}

// Compute two perpendicular vectors in the plane
func (p Plane) ComputeUVAxes() (Vector, Vector) {
	// Pick an arbitrary vector that is not parallel to the normal
	var arbitrary Vector
	if math.Abs(float64(p.normal.x)) < 0.9 {
		arbitrary = Vector{1, 0, 0} // Prefer X unless normal is mostly in X
	} else {
		arbitrary = Vector{0, 1, 0} // Otherwise, use Y
	}

	// Compute U axis as a perpendicular vector
	uAxis := arbitrary.CrossProduct(p.normal)
	uAxis.Normalize()

	// Compute V axis as the cross product of normal and U (ensuring orthogonality)
	vAxis := p.normal.CrossProduct(uAxis)
	vAxis.Normalize()

	return uAxis, vAxis
}
