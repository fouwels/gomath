package main

import (
	"math/cmplx"
)

func mandlebrot() error {

	points := []XY{}

	f := func(x complex128, c complex128) complex128 {
		return x*x + c
	}

	res := 100

	for cx := -1 * res; cx <= 1*res; cx++ {
		for ci := -1 * res; ci <= 1*res; ci++ {

			x := complex(0, 0)
			c := complex(float64(cx)/float64(res), float64(ci)/float64(res))

			for n := 0; n < 1000; n++ {
				x = f(x, c)
			}

			if cmplx.IsInf(x) || cmplx.IsNaN(x) {
				points = append(points, XY{X: real(c), Y: imag(c)})
				continue
			}

		}
	}

	return p(points, "mandlebrot", "x", "i")
}
