package color

import (
	"github.com/disintegration/imaging"
	"image"
	"image/color"
	"image/draw"
	"math/rand"
)

func MakeImage(colors []color.RGBA) image.Image {
	img := image.NewNRGBA(image.Rect(0, 0, 200, 200))

	rects := make([]image.Rectangle, 4)
	rects[0] = image.Rect(0, 0, 100, 100)
	rects[1] = image.Rect(0, 100, 100, 200)
	rects[2] = image.Rect(100, 0, 200, 100)
	rects[3] = image.Rect(100, 100, 200, 200)

	for i, v := range colors {
		draw.Draw(img, rects[i], &image.Uniform{v}, image.ZP, draw.Src)
	}

	return img
}

// Characteristic color of a Cluster
type Result struct {
	Size  int
	Color color.RGBA
}

type Results []Result

func (slice Results) Len() int {
	return len(slice)
}

func (slice Results) Less(i, j int) bool {
	return slice[i].Size < slice[j].Size
}

func (slice Results) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

// ...

func ColorMean(list []color.RGBA) color.RGBA {
	r, g, b, a := 0.0, 0.0, 0.0, 0.0
	n := float64(len(list))
	for _, v := range list {
		r += float64(v.R)
		g += float64(v.G)
		b += float64(v.B)
		a += float64(v.A)
	}
	return color.RGBA{
		uint8(r / n),
		uint8(g / n),
		uint8(b / n),
		uint8(a / n),
	}
}

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

//
//
//
func ScaleDown(img image.Image) image.Image {
	// TODO: Handle failures like this, when the image is Null
	return imaging.Resize(img, 200, 200, imaging.Lanczos)
}

//
//
//
func Groups(n int) [][]color.RGBA {
	groups := make([][]color.RGBA, n)
	for i, _ := range groups {
		groups[i] = make([]color.RGBA, 0)
	}
	return groups
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
