package main

import (
	"math/cmplx"
)

func julia() error {

	const cx = float64(0.6)
	const ci = float64(0.156)
	const res = 100

	c := complex(cx, ci)
	points := []XY{}

	f := func(x complex128, c complex128) complex128 {
		return x*x + c
	}

	for x := -1 * res; x <= 1*res; x++ {
		for xi := -1 * res; xi <= 1*res; xi++ {

			x := complex(float64(x)/float64(res), float64(xi)/float64(res))
			y := x

			for n := 0; n < 8; n++ {
				y = f(y, c)

				if cmplx.Abs(y) > 10 {
					points = append(points, XY{X: real(x), Y: imag(x)})
					break
				}

			}
		}
	}

	return p(points, "julia", "x", "i")
}
