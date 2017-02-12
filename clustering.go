package color

import (
	"image"
	"image/color"
	"math/rand"
	"sort"
	"time"
)

// How to compute the representative color of a cluster
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

// Return num_colors clusters for a given image
func KMeans(img image.Image, num_colors int) Results {
	results := make(map[color.RGBA]int)
	final := make(Results, 0)

	// Seed centroids as just random colors
	centroids := RandomColors(num_colors)
	groups := Groups(num_colors)
	pixels := ColorVector(img)

	// For number of iterations
	for i := 0; i < 200; i++ {

		// Create
		groups = Groups(num_colors)

		// For each pixel see which random group it belongs to
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

		// Update centroids
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

// Return the dominant colors of an image, computed by k-means
func DominantColors(i image.Image, num_colors int) (Results, error) {
	rand.Seed(time.Now().Unix())

	results := KMeans(i, num_colors)

	return results, nil
}
