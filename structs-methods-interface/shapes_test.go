package structs

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestAreas(t *testing.T) {
	testCases := []struct {
		desc  string
		shape Shape
		hasArea  float64
	}{
		{
			desc:  "Rectangle",
			shape: Rectangle{12, 6},
			hasArea:  72.0,
		},
		{
			desc:  "Circle",
			shape: Circle{10},
			hasArea:  314.1592653589793,
		},
		{
			desc:  "Triangle",
			shape: Triangle{12, 6},
			hasArea:  36.0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := tC.shape.Area()
			if got != tC.hasArea {
				t.Errorf("got %g want %g", got, tC.hasArea)
			}
		})
	}
}
