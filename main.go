package main

import (
	"fmt"
	"image/color"
	"math"
	"runtime"
	"sync"
	"time"
)

const (
	screenWidth  = 800
	screenHeight = 600
)

var camera = Vector{0, 0, 0}
var screenZ = (screenHeight + screenWidth) * 2
var translateX = screenWidth / 2
var translateY = screenHeight / 2

var light = Vector{-screenWidth * 10, -screenHeight * 10, -screenHeight}

var objs = []geometry{
	CheckerPlane{Plane: Plane{Vector{0, 0, float64(screenZ * 25)}, Vector{0, 0, 1}, color.RGBA{0, 0, 0, 0}}},
	Plane{Vector{0, float64(screenZ / 5), 0}, Vector{0, 1, 0}, color.RGBA{255, 175, 0, 0}},
	Sphere{Vector{float64(translateX) * .25, float64(translateY), float64(screenZ) * 4.5}, float64(screenZ / 10), color.RGBA{255, 100, 0, 0}},
	Sphere{Vector{float64(-translateX) * 1.5, float64(translateY), float64(screenZ) * 3}, float64(screenZ / 10), color.RGBA{0, 255, 0, 0}},
	Sphere{Vector{float64(translateX) * 1.5, float64(translateY), float64(screenZ) * 2.5}, float64(screenZ / 10), color.RGBA{128, 0, 128, 0}},
}

func main() {

	var pixels [screenWidth][screenHeight]color.RGBA
	numCores := runtime.NumCPU()
	fmt.Println("Using", numCores, "CPU cores")

	start := time.Now()
	colsPerBatch := screenWidth / numCores

	//for range 500 {
	var wg sync.WaitGroup
	for i := range numCores {
		startCol := i * colsPerBatch
		endCol := startCol + colsPerBatch
		if i == numCores-1 {
			endCol = screenWidth
		}

		wg.Add(1)
		go processBatch(startCol, endCol, &pixels, &wg)
	}

	wg.Wait()
	//}

	since1 := time.Since(start)
	println(since1.Seconds())

	// Draw it
	pixelSlice := make([][]color.RGBA, screenWidth)
	for i := range pixelSlice {
		pixelSlice[i] = pixels[i][:] // Convert row to slice
	}
	var drawer ScreenDrawer = Screen{}
	drawer.Draw(&pixelSlice)
}

func processBatch(startRow, endRow int, pixels *[screenWidth][screenHeight]color.RGBA, wg *sync.WaitGroup) {
	defer wg.Done()

	for x := startRow; x < endRow; x++ {
		for y := 0; y < screenHeight; y++ {
			processPixel(pixels, x, y)
		}
	}
}

func processPixel(pixels *[screenWidth][screenHeight]color.RGBA, x int, y int) {
	v := Vector{float64(x - translateX), float64(y - translateY), float64(screenZ)}
	v.Normalize()
	minDistance := math.Inf(1)
	var closestObj geometry
	var closestIntersection Vector
	hasIntersection := false
	// Find the intersection of the object that is closest to the camera
	// if any
	for _, obj := range objs {
		intersects, i1 := obj.IntersectsAt(camera, v)
		if intersects {
			hasIntersection = true
			distance := camera.Distance(i1)
			if distance < minDistance {
				closestObj = obj
				closestIntersection = i1
				minDistance = distance
			}
		}
	}

	if hasIntersection {
		c := doLight(closestIntersection, light, closestObj, objs)
		//c := obj.Color()
		pixels[x][y] = c
	}
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

func doLight(phit Vector, light Vector, geo geometry, geometries []geometry) color.RGBA {
	/*
			            Vec3f transmission = 1;
		                Vec3f lightDirection = spheres[i].center - phit;
		                lightDirection.normalize();
		                for (unsigned j = 0; j < spheres.size(); ++j) {
		                    if (i != j) {
		                        float t0, t1;
								// looking for blockers...
		                        if (spheres[j].intersect(phit + nhit * bias, lightDirection, t0, t1)) {
		                            transmission = 0;
		                            break;
		                        }
		                    }
		                }

		                surfaceColor += sphere->surfaceColor * transmission *
		                std::max(float(0), nhit.dot(lightDirection)) * spheres[i].emissionColor;
	*/

	shadowRay := light.Subtract(phit)
	shadowRay.Normalize()
	nhit := geo.GetSurfaceNormal(phit)
	nhit.Normalize()

	// are we blocked by another object?
	transmission := 1
	for _, x := range geometries {
		if x == geo {
			continue
		}

		intersects, _ := x.IntersectsAt(phit.Add(nhit), shadowRay)
		if intersects {
			transmission = 0
			break
		}
	}
	c := math.Max(float64(0), nhit.DotProduct(shadowRay))
	return MultiplyColor(geo.Color(phit), c*float64(transmission))
}
