package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"os"
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

	for k, _ := range colors {
		colors[k] = dominants[k].Color
	}

	i := MakeImage(colors)
	OutputImage("ok.jpg", i)
	OutputImage("no.jpg", ScaleDown(img))
	fmt.Println("Fin.")
}
