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
			v.X() - w.X(),
			v.Y() - w.Y(),
			v.Z() - w.Z(),
		},
	}
}

func (v Vector) Multiply(scalar float32) Vector {
	return Vector{
		val: [3]float32{
			v.X() * scalar,
			v.Y() * scalar,
			v.Z() * scalar,
		},
	}
}

func (v Vector) Add(w Vector) Vector {
	return Vector{
		val: [3]float32{
			v.X() + w.X(),
			v.Y() + w.Y(),
			v.Z() + w.Z(),
		},
	}
}

func (v Vector) DotProduct(other Vector) float32 {
	return v.X()*other.X() + v.Y()*other.Y() + v.Z()*other.Z()
}
