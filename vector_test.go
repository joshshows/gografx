package main

import (
	"testing"
)

func TestVectorNormalize(t *testing.T) {
	v := Vector{val: [3]float32{21.4, 14.22, 71.43}}
	v.Normalize()

	expected := [3]float32{0.28191072, 0.18732572, 0.9409758}
	if expected != v.val {
		t.Fatalf(`Arrays did not match %v %v`, expected, v.val)
	}
}
