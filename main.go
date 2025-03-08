package main

import (
	"image/color"
	"math"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

func main() {
	// camera is at 0,0,0 in world units.
	camera := Vector{0, 0, 0}
	// view screen is some distance away based on resolution.
	screenZ := (screenHeight + screenWidth) * 2
	translateX := screenWidth / 2
	translateY := screenHeight / 2

	light := Sphere{
		center: Vector{-1000, 1000, float64(screenZ)},
		radius: float64(screenZ / 5),
	}

	/*
		sphere := Sphere{
			center: Vector{0, 0, float64(screenZ * 2)},
			radius: float64(screenZ / 10),
		}*/

	objs := [...]Sphere{
		{Vector{0, 0, float64(screenZ * 2)}, float64(screenZ / 10), color.RGBA{0, 0, 255, 0}},
		{Vector{float64(translateX), float64(translateY), float64(screenZ * 4)}, float64(screenZ / 10), color.RGBA{255, 0, 0, 0}},
		{Vector{float64(-translateX), float64(-translateY), float64(screenZ * 4)}, float64(screenZ / 5), color.RGBA{0, 255, 0, 0}},
	}

	pixels := make([][]color.RGBA, screenWidth)
	for i := range pixels {
		pixels[i] = make([]color.RGBA, screenHeight)
	}

	for x := range screenWidth {
		for y := range screenHeight {
			v := Vector{float64(x - translateX), float64(y - translateY), float64(screenZ)}
			v.Normalize()
			//intersects := v.Intersects(sphere)
			for _, obj := range objs {
				intersects, i1, _ := obj.IntersectsAt(camera, v)
				if intersects {
					pixels[x][y] = doLight(i1, light, obj) //color.RGBA{0, 0, uint8(255 * c), 255}
				} /* else {
					pixels[x][y] = color.RGBA{0, 0, 0, 255}
				}*/
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

func doLight(phit Vector, light Sphere, object Sphere) color.RGBA {
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

	shadowRay := light.center.Subtract(phit)
	shadowRay.Normalize()
	nhit := phit.Subtract(object.center)
	nhit.Normalize()
	c := math.Max(float64(0), nhit.DotProduct(shadowRay))
	return multiplyColor(object.color, c)
}

func multiplyColor(c color.RGBA, x float64) color.RGBA {
	fac := [3]float64{float64(c.R), float64(c.G), float64(c.B)}
	return color.RGBA{uint8(fac[0] * x), uint8(fac[1] * x), uint8(fac[2] * x), 255}
}
