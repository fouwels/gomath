package main

func lorenz() error {

	const time = 100
	const step = 0.01

	points := []XYZ{}

	x := float64(0.1)
	y := float64(0)
	z := float64(0)

	for t := 0 * step; t <= time; t += step {
		x, y, z = l(step, x, y, z)

		points = append(points, XYZ{
			X: x,
			Y: y,
			Z: z,
		})

	}

	return p3(points, "lorenz", "x", "y", "z")
}

func l(t, x, y, z float64) (float64, float64, float64) {

	const q = 10.0
	const p = 30.0
	const B = 8.0 / 3.0

	dx := t * (q * (y - x))
	dy := t * (x*(p-z) - y)
	dz := t * (x*y - B*z)

	return x + dx, y + dy, z + dz
}
