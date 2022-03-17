package main

import (
	"fmt"
	"image/color"
	"log"
	"os"

	"github.com/Arafatk/glot"
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

func p3(points []plotter.XYZ, name string, x string, y string, z string) error {

	ps := make([][]float64, 3)

	for _, v := range points {
		ps[0] = append(ps[0], v.X)
		ps[1] = append(ps[1], v.Y)
		ps[2] = append(ps[2], v.Z)

		//log.Printf("%v %v %v", v.X, v.Y, v.Z)
	}

	//plot 3D

	plot, err := glot.NewPlot(3, false, false)
	if err != nil {
		return fmt.Errorf("failed to create plot: %v", err)
	}

	err = plot.SetTitle(name)
	if err != nil {
		return fmt.Errorf("failed to set title: %v", err)
	}
	err = plot.SetXLabel(x)
	if err != nil {
		return fmt.Errorf("failed to set x label: %v", err)
	}
	err = plot.SetYLabel(y)
	if err != nil {
		return fmt.Errorf("failed to set y label: %v", err)
	}
	err = plot.SetZLabel(z)
	if err != nil {
		return fmt.Errorf("failed to set z label: %v", err)
	}

	//plot 3D
	err = plot.AddPointGroup("points", "dots", [][]float64{
		ps[0], ps[1], ps[2],
	})
	if err != nil {
		return fmt.Errorf("failed to add point group: %v", err)
	}
	err = plot.SavePlot("points-" + name + ".png")
	if err != nil {
		return fmt.Errorf("failed to save plot: %v", err)
	}

	//plot 2D
	plot, err = glot.NewPlot(2, false, false)
	if err != nil {
		return fmt.Errorf("failed to create plot: %v", err)
	}

	err = plot.AddPointGroup("x", "dots", ps[0])
	if err != nil {
		return fmt.Errorf("failed to add point group: %v", err)
	}
	err = plot.AddPointGroup("y", "dots", ps[1])
	if err != nil {
		return fmt.Errorf("failed to add point group: %v", err)
	}
	err = plot.AddPointGroup("z", "dots", ps[2])
	if err != nil {
		return fmt.Errorf("failed to add point group: %v", err)
	}

	err = plot.SetTitle(name + "-2d")
	if err != nil {
		return fmt.Errorf("failed to set title: %v", err)
	}

	err = plot.SetXLabel("t")
	if err != nil {
		return fmt.Errorf("failed to set x label: %v", err)
	}
	err = plot.SetYLabel("value")
	if err != nil {
		return fmt.Errorf("failed to set y label: %v", err)
	}
	err = plot.SetLogscale("y", 10)
	if err != nil {
		return fmt.Errorf("failed to set y log: %v", err)
	}

	err = plot.SavePlot("points-" + name + "-2d.png")
	if err != nil {
		return fmt.Errorf("failed to save plot: %v", err)
	}

	return nil
}
