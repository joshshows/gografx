package main

import (
	"math"
)

type Vector struct {
	val [3]float32
}

func (v Vector) X() float32 {
	return v.val[0]
}

func (v Vector) Y() float32 {
	return v.val[1]
}
func (v Vector) Z() float32 {
	return v.val[2]
}

func (v *Vector) Normalize() {
	length := float32(math.Sqrt(float64(v.val[0]*v.val[0] + v.val[1]*v.val[1] + v.val[2]*v.val[2])))
	for i := range v.val {
		v.val[i] = v.val[i] / length
	}
}

func (v Vector) Subtract(w Vector) Vector {
	return Vector{
		val: [3]float32{
			v.val[0] - w.val[0],
			v.val[1] - w.val[1],
			v.val[2] - w.val[2],
		},
	}
}

func (v Vector) DotProduct(other Vector) float32 {
	return v.val[0]*other.val[0] + v.val[1]*other.val[1] + v.val[2]*other.val[2]
}

func (v Vector) Intersects(sphere Sphere) bool {
	origin := Vector{val: [3]float32{0, 0, 0}}
	L := origin.Subtract(sphere.center)
	b := 2.0 * v.DotProduct(L)
	c := L.DotProduct(L) - sphere.radius*sphere.radius

	discriminant := b*b - 4*c
	return discriminant >= 0
}
