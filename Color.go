package main

import (
	"github.com/disintegration/imaging"
	"image"
	"image/color"
	"image/draw"
	"math"
	"math/rand"
	"sort"
	"time"
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

// Results object

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

func Mean(list []float64) float64 {
	sum := 0.0
	for _, v := range list {
		sum += v
	}
	return sum / float64(len(list))
}

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
//
//
func Kmeans(img image.Image, num_colors int) Results {
	results := make(map[color.RGBA]int)
	final := make(Results, 0)

	centroids := RandomColors(num_colors)
	groups := Groups(num_colors)
	pixels := ColorVector(img)

	for i := 0; i < 200; i++ {

		// Create
		groups = Groups(num_colors)

		//
		for _, pixel := range pixels {
			// Compute difference
			differences := make([]float64, num_colors)

			// Differences
			for k, _ := range differences {
				differences[k] = Distance(pixel, centroids[k])
			}

			// Figure out the closest cluster
			index, _ := ArgMin(differences)

			// Add to the most apropriate cluster
			groups[index] = append(groups[index], pixel)
		}

		centroids_old := make([]color.RGBA, len(groups))
		copy(centroids_old[:], centroids)
		centroid_diff := make([]float64, num_colors)

		// Put each pixel into a cluster
		for j, _ := range groups {
			centroids[j] = ColorMean(groups[j])
			centroid_diff[j] = Distance(centroids[j], centroids_old[j])
		}

		_, max_diff := ArgMax(centroid_diff)

		if max_diff < 1 {
			break
		}
	}

	// Measure the size of each cluster
	for i, v := range centroids {
		results[v] = len(groups[i])
	}

	for color, count := range results {
		final = append(final, Result{count, color})
	}

	sort.Sort(final)

	return final
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
	r := math.Abs(float64(color1.R) - float64(color2.R))
	g := math.Abs(float64(color1.G) - float64(color2.G))
	b := math.Abs(float64(color1.B) - float64(color2.B))
	_, max := ArgMax([]float64{r, g, b})
	return max
}

// Images
func DominantColors(i image.Image, num_colors int) ([]Result, error) {
	rand.Seed(time.Now().Unix())

	results := Kmeans(i, num_colors)

	return results, nil
}
