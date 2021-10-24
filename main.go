package main

import (
	"image/color"
	"log"
	"os"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	err := run()
	log.Printf("Exit with err: %v", err)
	os.Exit(1)
}
func run() error {
	return bifurcation()
}

func bifurcation() error {

	points := []plotter.XY{}

	res := 100

	f := func(x complex128, r complex128) complex128 {
		return r * x * (1 - x)
	}

	for ri := 0; ri <= 4*res; ri++ {

		r := complex(float64(ri)/float64(res), 0)
		x := complex(0.1, 0)

		for n := 0; n < 1000; n++ {
			x = f(x, r)

			if n > 900 {
				points = append(points, plotter.XY{X: real(r), Y: real(x)})
			}
		}
	}

	return p(points, "bifurcation", "r", "x")
}

func p(points []plotter.XY, name string, x string, y string) error {
	p := plot.New()

	p.Title.Text = name
	p.X.Label.Text = x
	p.Y.Label.Text = y

	s, err := plotter.NewScatter(plotter.XYs(points))
	if err != nil {
		return err
	}
	s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}
	s.Radius = 1

	p.Add(s)

	// Save the plot to a PNG file.
	err = p.Save(20*vg.Centimeter, 20*vg.Centimeter, "points-"+name+".png")
	if err != nil {
		return err
	}

	return nil
}
