package main

import (
	"github.com/disintegration/imaging"
	"image"
	"image/color"
	"math"
	"math/rand"
)

func ColorVector(img image.Image) []color.RGBA {
	scaled_down_image := ScaleDown(img)
	bounds := scaled_down_image.Bounds()
	colors := []color.RGBA{}
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			col := img.At(x, y)
			r, g, b, a := col.RGBA()
			colors = append(colors, color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)})
		}
	}

	return colors
}

// Returns a List of Centroidsand their Clusters
func Closest(img image.Image, num_colors int) (color.RGBA, []color.RGBA) {
	// A+
	groups := make([][]color.RGBA, num_colors)

	return color.RGBA{}, groups[0]
}

func ScaleDown(img image.Image) image.Image {
	return imaging.Resize(img, 200, 200, imaging.Lanczos)
}

func Kmeans(img image.Image, num_colors int) []color.RGBA {
	centroids := make([]color.RGBA, 10)
	for i := 0; i < 100; i++ {
		_, centroids = Closest(img, num_colors)
	}

	return centroids
}

//
func RandomColor() color.RGBA {
	return color.RGBA{
		uint8(rand.Intn(255)),
		uint8(rand.Intn(255)),
		uint8(rand.Intn(255)),
		255,
	}
}

// Images
func RandomColors(num_colors int) []color.RGBA {
	palette := make([]color.RGBA, num_colors)
	for i, _ := range palette {
		palette[i] = RandomColor()
	}
	return palette
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
