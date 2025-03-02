package main

import (
	"image/color"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

func main() {
	// camera is at 0,0,0 in world units.
	camera := Vector{val: [3]float32{0, 0, 0}}
	// view screen is some distance away based on resolution.
	screenZ := (screenHeight + screenWidth) * 2
	translateX := screenWidth / 2
	translateY := screenHeight / 2

	sphere := Sphere{
		center: Vector{val: [3]float32{0, 0, float32(screenZ * 2)}},
		radius: float32(screenZ / 10),
	}

	pixels := make([][]color.RGBA, screenWidth)
	for i := range pixels {
		pixels[i] = make([]color.RGBA, screenHeight)
	}

	for x := range screenWidth {
		for y := range screenHeight {
			v := Vector{val: [3]float32{float32(x - translateX), float32(y - translateY), float32(screenZ)}}
			v.Normalize()
			//intersects := v.Intersects(sphere)
			intersects, _, _ := sphere.IntersectsAt(camera, v)
			if intersects {
				pixels[x][y] = color.RGBA{0, 0, 255, 255}
			} else {
				pixels[x][y] = color.RGBA{0, 0, 0, 255}
			}
		}
	}

	// Draw it
	var drawer ScreenDrawer = Screen{}
	drawer.Draw(&pixels)

}

func doSomeStuff() {

	// Create the screen array.
	pixels := make([][]color.RGBA, screenWidth)
	for i := range pixels {
		pixels[i] = make([]color.RGBA, screenHeight)
	}

	// Populate with a gradient
	for x := range screenWidth {
		for y := range screenHeight {
			pixels[x][y] = color.RGBA{uint8(x), uint8(y), 255, 255}
		}
	}

	// Draw it
	var drawer ScreenDrawer = Screen{}
	drawer.Draw(&pixels)
}
