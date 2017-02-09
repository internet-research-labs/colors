package color

import (
	"fmt"
	"image/color"
	"testing"
)

func TestLuminance(t *testing.T) {
	a := color.RGBA{0, 0, 0, 255}
	if v := Luminance(a); v != 0.0 {
		t.Fail()
	}

	b := color.RGBA{20, 20, 10, 255}
	if v := Luminance(b); v != 19.299999999999997 {
		t.Fail()
	}

}

func TestDistance(t *testing.T) {
	a := color.RGBA{0, 0, 0, 255}
	b := color.RGBA{0, 0, 0, 255}
	d := color.RGBA{2, 10, 23, 255}

	if Distance(a, b) != 0.0 {
		t.Fail()
	}

	if Norm(a) != 0.0 {
		t.Fail()
	}

	if v := Distance(d, a); v != 26.79912125425011 {
		fmt.Println(v)
		t.Fail()
	}

	if Norm(d) != 26.79912125425011 {
		t.Fail()
	}
}

func _() {
	fmt.Println("x_x")
}
