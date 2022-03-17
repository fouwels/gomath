package main

import (
	"math"

	"gonum.org/v1/plot/plotter"
)

func lorenz() error {

	const time = 1000
	const step = 0.01

	points := []plotter.XYZ{}

	x := float64(0.1)
	y := float64(0)
	z := float64(0)

	for t := 0 * step; t <= time; t += step {
		x, y, z = l(step, x, y, z)

		if !math.IsInf(x, 0) && !math.IsNaN(x) && !math.IsInf(y, 0) && !math.IsNaN(y) && !math.IsInf(z, 0) && !math.IsNaN(z) {
			points = append(points, plotter.XYZ{
				X: x,
				Y: y,
				Z: z,
			})
			continue
		}
	}

	return p3(points, "lorenz", "x", "y", "z")
}

func l(t, x, y, z float64) (float64, float64, float64) {

	const q = 10.0
	const p = 28.0
	const B = 8.0 / 3.0

	dx := t * (q * (y - x))
	dy := t * (x*(p-z) - y)
	dz := t * (x*y - B*z)

	return x + dx, y + dy, z + dz
}
