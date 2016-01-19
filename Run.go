package main

import (
	"encoding/csv"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"
	"path"
	"time"
)

func ColorToHex(c color.RGBA) string {
	return fmt.Sprintf("%02X%02X%02X", c.R, c.G, c.B)
}

// Load Image
func LoadImage(filename string) image.Image {
	file, _ := os.Open(filename)
	defer file.Close()
	img, _ := jpeg.Decode(file)
	return img
}

func OutputImage(filename string, img image.Image) {
	dst, _ := os.Create(filename)
	defer dst.Close()
	jpeg.Encode(dst, img, &jpeg.Options{jpeg.DefaultQuality})
}

func main() {

	now := time.Now().Unix()
	filename := fmt.Sprintf("%d.csv", now)

	fileWriter, _ := os.Create(filename)
	defer fileWriter.Close()

	log.Printf("Wrinting %s", filename)

	// ...
	csvWriter := csv.NewWriter(fileWriter)
	record := []string{"okay"}
	csvWriter.Write(record)

	// filename := os.Args[1]
	// img := LoadImage(filename)
	// dominants := DominantColors(img, 4)

	colors := make([]color.RGBA, 4)

	// ...
	dir_name := fmt.Sprintf("%d", now)

	// ...
	os.MkdirAll(dir_name, 0777)

	// Create connection
	s := NewInstagramConnection(
		"1553cfcee2b74ad9ba8f75b0a278b6ac",
		"picoftheday",
	)

	// s.Get()
	s.Start()

	// suffix for images
	count := 0

	for {
		actual_image := <-s.Images()
		actual_dst := path.Join(".", dir_name, fmt.Sprintf("%05d_actual.jpg", count))
		dominants := DominantColors(actual_image, 4)

		// TODO: Fix this failure condition
		if len(dominants) != 4 {
			continue
		}

		// Writing to CSV
		csvWriter.Write([]string{
			fmt.Sprintf(actual_dst),
			ColorToHex(dominants[0].Color),
			ColorToHex(dominants[1].Color),
			ColorToHex(dominants[2].Color),
			ColorToHex(dominants[3].Color),
			fmt.Sprintf("%d", dominants[0].Size),
			fmt.Sprintf("%d", dominants[1].Size),
			fmt.Sprintf("%d", dominants[2].Size),
			fmt.Sprintf("%d", dominants[3].Size),
		})

		// Flushing
		csvWriter.Flush()

		for k, _ := range colors {
			colors[k] = dominants[k].Color
		}

		computed_dst := path.Join(".", dir_name, fmt.Sprintf("%05d_computed.jpg", count))

		// Outputting images
		OutputImage(actual_dst, actual_image)
		OutputImage(computed_dst, MakeImage(colors))

		// A+
		count += 1
	}
}
