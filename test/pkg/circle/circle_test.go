package circle

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	circle := Circle{Radius: 1}
	got := circle.Area()
	expected := math.Pi

	if got != expected {
		t.Errorf("got %f, expected %f", got, expected)
	}
}
