package main

import (
	"reflect"
	"testing"
)

func TestVectorNormalize(t *testing.T) {
	v := Vector{21.4, 14.22, 71.43}
	v.Normalize()

	expected := Vector{0.2819107073490415, 0.1873257130141762, 0.9409757862589737}
	if !reflect.DeepEqual(expected, v) {
		t.Fatalf(`Arrays did not match %v %v`, expected, v)
	}
}
