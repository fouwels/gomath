package main

import (
	"log"
	"testing"
)

func TestBifurcation(t *testing.T) {
	err := bifurcation()
	log.Printf("%v", err)
}

func TestJulia(t *testing.T) {
	err := julia()
	log.Printf("%v", err)
}

func TestMandlebrot(t *testing.T) {
	err := mandlebrot()
	log.Printf("%v", err)
}

// func TestLorenz(t *testing.T) {
// 	err := lorenz()
// 	log.Printf("%v", err)
// }
