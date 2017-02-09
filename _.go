package color

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
