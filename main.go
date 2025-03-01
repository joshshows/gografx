package main

import (
	"image/color"
)

const (
	screenWidth  = 256
	screenHeight = 256
)

func main() {

	// Create the screen array.
	pixels := make([][]color.RGBA, screenWidth)
	for i := range pixels {
		pixels[i] = make([]color.RGBA, screenHeight)
	}

	// Populate with a gradient
	for x := 0; x < screenWidth; x++ {
		for y := 0; y < screenHeight; y++ {
			pixels[x][y] = color.RGBA{uint8(x), uint8(y), 255, 255}
		}
	}

	// Draw it
	var drawer ScreenDrawer = Screen{}
	drawer.Draw(pixels)
}
