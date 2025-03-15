package main

import (
	"image/color"
	"math"
)

const (
	screenWidth  = 800
	screenHeight = 600
)

func main() {
	camera := Vector{0, 0, 0}
	// build the scene relative to screen resolution
	screenZ := (screenHeight + screenWidth) * 2
	translateX := screenWidth / 2
	translateY := screenHeight / 2

	light := Vector{-screenWidth * 10, -screenHeight * 10, -screenHeight}

	objs := []geometry{
		//Plane{Vector{0, 0, float64(screenZ * 25)}, Vector{0, 0, 1}, color.RGBA{255, 255, 255, 0}},
		Plane{Vector{0, float64(screenZ / 5), 0}, Vector{0, 1, 0}, color.RGBA{0, 127, 255, 0}},
		Sphere{Vector{float64(translateX) * .25, float64(translateY), float64(screenZ * 4)}, float64(screenZ / 10), color.RGBA{255, 165, 0, 0}},
		Sphere{Vector{float64(-translateX) * 1.5, float64(translateY), float64(screenZ) * 2.5}, float64(screenZ / 10), color.RGBA{0, 255, 0, 0}},
		Sphere{Vector{float64(translateX) * 1.5, float64(translateY), float64(screenZ * 2)}, float64(screenZ / 10), color.RGBA{128, 0, 128, 0}},
	}

	pixels := make([][]color.RGBA, screenWidth)
	for i := range pixels {
		pixels[i] = make([]color.RGBA, screenHeight)
	}

	for x := range screenWidth {
		for y := range screenHeight {
			v := Vector{float64(x - translateX), float64(y - translateY), float64(screenZ)}
			v.Normalize()
			for _, obj := range objs {
				intersects, i1 := obj.IntersectsAt(camera, v)
				if intersects {
					c := doLight(i1, light, obj, objs)
					//c := obj.Color()
					pixels[x][y] = c
				}
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
