package color

import (
	"fmt"
	"testing"
)

func TestArgMin(t *testing.T) {
	if i, _ := ArgMin(nil); i != -1 {
		t.Fail()
	}
	if _, v := ArgMin([]float64{1.0, 4.0}); v != 1.0 {
		fmt.Println(v)
		t.Fail()
	}
	if i, _ := ArgMin([]float64{1.0, 4.0}); i != 0 {
		t.Fail()
	}
}

func TestArgMax(t *testing.T) {
	if _, v := ArgMax([]float64{1.0, 4.0}); v != 4.0 {
		fmt.Println(v)
		t.Fail()
	}
	if i, _ := ArgMax([]float64{1.0, 4.0}); i != 1 {
		t.Fail()
	}
}

func TestMean(t *testing.T) {
	if v := Mean([]float64{1.0, 2.0}); v != 1.5 {
		t.Fail()
	}
}
