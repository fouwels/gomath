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

func p3(points []plotter.XYZ, name string, x string, y string, z string) error {
	p := plot.New()

	p.Title.Text = name
	p.X.Label.Text = x
	p.Y.Label.Text = y
	p.Z.Label.Text = z

	s, err := plotter.New(plotter.XYZs(points))
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
