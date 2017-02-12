package color

import (
	"fmt"
	"image/color"
	"math"
)

// Return the ArgMin of a slice of floats
func ArgMin(list []float64) (int, float64) {
	argmin := -1
	min := math.Inf(+1)
	for i, v := range list {
		if v < min {
			min = v
			argmin = i
		}
	}
	return argmin, min
}

// Return the ArgMax of a slice of floats
func ArgMax(list []float64) (int, float64) {
	argmax := -1
	max := math.Inf(-1)
	for i, v := range list {
		if v > max {
			max = v
			argmax = i
		}
	}
	return argmax, max
}

// Return the arithmetic mean of a set of floats
func Mean(list []float64) float64 {
	sum := 0.0
	for _, v := range list {
		sum += v
	}
	return sum / float64(len(list))
}

// Return a hex-string from a color
func ColorToHex(c color.RGBA) string {
	return fmt.Sprintf("%02X%02X%02X", c.R, c.G, c.B)
}
