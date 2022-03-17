package main

import (
	"log"
	"math"

	"gonum.org/v1/plot/plotter"
)

func lorenz() error {

	const time = 100
	const step = 1

	points := []plotter.XYZ{}

	x := float64(1)
	y := float64(1)
	z := float64(1)

	for t := 0 * step; t <= time; t += step {
		log.Printf("%v: %v %v %v", t, x, y, z)

		x, y, z = l(step, x, y, z)

		if !math.IsInf(x, 0) && !math.IsNaN(x) && !math.IsInf(y, 0) && !math.IsNaN(y) && !math.IsInf(z, 0) && !math.IsNaN(z) {
			points = append(points, plotter.XYZ{x, y, z})
			continue
		}
	}

	//lorenz = @(t,x) [];

	return p3(points, "lorenz", "x", "y", "z")
}

func l(t, x, y, z float64) (dx, dy, dz float64) {

	const q = 10
	const p = 28
	const B = 8 / 3

	dx = t * (q * (y - x))
	dy = t * (x*(p-z) - y)
	dz = t * (z*y - B*z)

	return
}
