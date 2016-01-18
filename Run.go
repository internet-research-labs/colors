package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"os"
	"path"
	"time"
)

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
	filename := os.Args[1]
	img := LoadImage(filename)
	dominants := DominantColors(img, 4)
	colors := make([]color.RGBA, 4)

	fmt.Println(dominants)

	// ...
	dir_name := fmt.Sprintf("%d", time.Now().Unix())

	// ...
	os.MkdirAll(dir_name, 0777)

	for k, _ := range colors {
		colors[k] = dominants[k].Color
	}

	i := MakeImage(colors)
	OutputImage("ok.jpg", i)
	OutputImage("no.jpg", ScaleDown(img))
	fmt.Println("Fin.")

	s := NewInstagramConnection(
		"1553cfcee2b74ad9ba8f75b0a278b6ac",
		"yolo",
	)

	s.Get()

	// suffix for images
	count := 0

	for {
		actual_image := <-s.Images()
		actual_dst := path.Join(".", dir_name, fmt.Sprintf("%05d_actual.jpg", count))
		dominants := DominantColors(actual_image, 4)
		for k, _ := range colors {
			colors[k] = dominants[k].Color
		}
		computed_dst := path.Join(".", dir_name, fmt.Sprintf("%05d_computed.jpg", count))
		count += 1
		fmt.Println(fmt.Sprintf("Writing to: '%s'", actual_dst))
		OutputImage(actual_dst, actual_image)
		OutputImage(computed_dst, MakeImage(colors))
	}
}
