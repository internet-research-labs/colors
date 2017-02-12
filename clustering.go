package color

import (
	"image"
	"image/color"
	"math/rand"
	"sort"
	"time"
)

//
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
