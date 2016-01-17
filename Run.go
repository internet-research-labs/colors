package main

import (
	"fmt"
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

func main() {
	img := LoadImage("images/neo-tokyo.jpg")
	fmt.Println(DominantColors(img, 5))
}
