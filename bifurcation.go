package main

func bifurcation() error {

	const res = 100
	const rmax = 4

	points := []XY{}

	f := func(x complex128, r complex128) complex128 {
		return r * x * (1 - x)
	}

	for ri := 0; ri <= rmax*res; ri++ {

		r := complex(float64(ri)/float64(res), 0)
		x := complex(0.1, 0)

		for n := 0; n < 1000; n++ {
			x = f(x, r)

			if n > 900 {
				points = append(points, XY{X: real(r), Y: real(x)})
			}
		}
	}

	return p(points, "bifurcation", "r", "x")
}
