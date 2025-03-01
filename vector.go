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
