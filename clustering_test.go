package color

import (
	"fmt"
	"testing"
)

func TestKMeans(t *testing.T) {
	runaway_yeezy := LoadImage("images/Runawayyeezy.jpg")
	runaway_palette := KMeans(runaway_yeezy, 4)
	fmt.Println(runaway_palette)
}

func _() {
	fmt.Println("_")
}
