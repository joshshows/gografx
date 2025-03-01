package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type ScreenDrawer interface {
	Draw(pixels *[][]color.RGBA)
}

type Screen struct{}

func (s Screen) Draw(pixels *[][]color.RGBA) {
	width := len(*pixels)
	height := len((*pixels)[0])
	ebiten.SetWindowSize(width*2, height*2)
	ebiten.SetWindowTitle("2D Pixel Array")

	game := &Game{
		pixels: pixels,
	}
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

type Game struct {
	pixels *[][]color.RGBA
}

func (g *Game) Update() error {
	// Update pixel colors dynamically if needed
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	width := len(*g.pixels)
	height := len((*g.pixels)[0])

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			screen.Set(x, y, (*g.pixels)[x][y])
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return len(*g.pixels), len((*g.pixels)[0])
}
