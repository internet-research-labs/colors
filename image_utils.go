package color

import (
	"image"
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
