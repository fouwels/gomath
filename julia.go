package main

import (
	"math/cmplx"

	"gonum.org/v1/plot/plotter"
)

func julia() error {
	points := []plotter.XY{}

	f := func(x complex128, c complex128) complex128 {
		return x*x + c
	}

	cx := float64(0.8)
	ci := float64(0.156)
	c := complex(cx, ci)

	res := 100

	for x := -1 * res; x <= 1*res; x++ {
		for xi := -1 * res; xi <= 1*res; xi++ {

			x := complex(float64(x)/float64(res), float64(xi)/float64(res))
			y := x

			for n := 0; n < 8; n++ {
				y = f(y, c)

				if cmplx.Abs(y) > 10 {
					points = append(points, plotter.XY{X: real(x), Y: imag(x)})
					break
				}

			}
		}
	}

	return p(points, "julia", "x", "i")
}
