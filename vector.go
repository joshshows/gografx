package main

import (
	"math"
)

type Vector struct {
	x float64
	y float64
	z float64
}

func (v *Vector) Normalize() {
	length := math.Sqrt(v.x*v.x + v.y*v.y + v.z*v.z)
	v.x = v.x / length
	v.y = v.y / length
	v.z = v.z / length
}

func (v Vector) Subtract(w Vector) Vector {
	return Vector{
		x: v.x - w.x,
		y: v.y - w.y,
		z: v.z - w.z,
	}
}

func (v Vector) Multiply(scalar float64) Vector {
	return Vector{
		v.x * scalar,
		v.y * scalar,
		v.z * scalar,
	}
}

func (v Vector) Add(w Vector) Vector {
	return Vector{
		v.x + w.x,
		v.y + w.y,
		v.z + w.z,
	}
}

func (v Vector) DotProduct(other Vector) float64 {
	return v.x*other.x + v.y*other.y + v.z*other.z
}

func (v Vector) CrossProduct(b Vector) Vector {
	return Vector{
		x: v.y*b.z - v.z*b.y,
		y: v.z*b.x - v.x*b.z,
		z: v.x*b.y - v.y*b.x,
	}
}

func (v Vector) Distance(v2 Vector) float64 {
	//d = √((x2 - x1)² + (y2 - y1)² + (z2 - z1)²)
	return math.Sqrt(math.Pow(v2.x-v.x, 2) + math.Pow(v2.y-v.y, 2) + math.Pow(v2.z-v.z, 2))
}
