package color

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

func OutputImage(filename string, img image.Image) {
	dst, _ := os.Create(filename)
	defer dst.Close()
	jpeg.Encode(dst, img, &jpeg.Options{jpeg.DefaultQuality})
}

// Main
func main() {

	// Setup variables
	now := time.Now().Unix()
	filename := fmt.Sprintf("%d.csv", now)

	// Open file
	fileWriter, err := os.Create(filename)

	if err != nil {
		log.Panic(fmt.Sprintf("Failed to open file \"%s\"", filename))
	}

	// Defer close
	defer fileWriter.Close()

	log.Printf("Wrinting %s", filename)

	// ...
	csvWriter := csv.NewWriter(fileWriter)
	record := []string{"okay"}
	csvWriter.Write(record)

	colors := make([]color.RGBA, 4)

	// ...
	dir_name := fmt.Sprintf("%d", now)

	// ...
	os.MkdirAll(dir_name, 0777)

	// Create connection
	s := NewInstagramConnection(
		"1553cfcee2b74ad9ba8f75b0a278b6ac",
		"blizzard2016",
	)

	// s.Get()
	s.Start()

	// suffix for images
	count := 0

	for {
		actual_image := <-s.Images()
		file_name := "%05d_actual.jpg"
		actual_dst := path.Join(".", dir_name, fmt.Sprintf(file_name, count))
		if actual_image == nil {
			log.Printf("Skipping file \"%s\".", file_name)
			continue
		}
		dominants, err := DominantColors(actual_image, 4)

		if err != nil {
			continue
		}

		// TODO: Fix this failure condition
		if len(dominants) != 4 {
			continue
		}

		// Writing to CSV
		csvWriter.Write([]string{
			// Timestamp
			string(time.Now().Unix()),
			// Filename
			fmt.Sprintf(actual_dst),
			//
			// Dominant Colors 1-4
			ColorToHex(dominants[0].Color),
			ColorToHex(dominants[1].Color),
			ColorToHex(dominants[2].Color),
			ColorToHex(dominants[3].Color),
			// Frequency of Dominant Color 1-4
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
