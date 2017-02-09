package color

import (
	"image/color"
	"math"
)

// Return the Luminance value of a color
func Luminance(c color.RGBA) float64 {
	r := 0.21 * float64(c.R)
	g := 0.72 * float64(c.G)
	b := 0.07 * float64(c.B)
	return r + g + b
}

/*
Return the Euclidean Distance between Two Colors

Args:
	lhs color as a vector
	rhs color as a vector
Return:
	Distance
*/
func Distance(lhs, rhs color.RGBA) float64 {
	r := float64(lhs.R) - float64(rhs.R)
	g := float64(lhs.G) - float64(rhs.G)
	b := float64(lhs.B) - float64(rhs.B)
	l := Luminance(lhs) - Luminance(rhs)
	return math.Sqrt(r*r + g*g + b*b + l*l)
}

// Return the magnitude (norm) of a color
func Norm(c color.RGBA) float64 {
	return Distance(c, color.RGBA{0, 0, 0, 255})
}
