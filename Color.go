package main

import (
	"image"
	"image/color"
	"math"
	"math/rand"
)

//
func RandomColor() color.RGBA {
	return color.RGBA{
		uint8(rand.Intn(255)),
		uint8(rand.Intn(255)),
		uint8(rand.Intn(255)),
		255,
	}
}

// Euclidean Distance
func Distance(color1 color.RGBA, color2 color.RGBA) float64 {
	r := float64(color1.R - color2.R)
	g := float64(color1.G - color2.G)
	b := float64(color1.B - color2.B)
	return math.Sqrt(r*r + g*g + b*b)
}

// Images
func DominantColors(i image.Image, num_colors int) []color.RGBA {
	palette := make([]color.RGBA, num_colors)

	for i, _ := range palette {
		palette[i] = RandomColor()
	}

	return palette
}
