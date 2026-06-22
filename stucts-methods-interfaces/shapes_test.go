package stuctsmethodsinterfaces

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rect := Rectangle{10, 15}
	got := Perimeter(rect)
	expect := 50.0

	if got != expect {
		t.Errorf(
			"\n - got: %g \n - expected: %g",
			got, expect,
		)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, expect float64) {
		t.Helper()
		got := shape.Area()

		if got != expect {
			t.Errorf(
				"\n - got: %g \n - expected: %g",
				got, expect,
			)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rect := Rectangle{10, 15}
		checkArea(t, rect, 25.0)
	})
	t.Run("circles", func(t *testing.T) {
		circle := Circle{8}
		checkArea(t, circle, 201.06192982974676)
	})
}

// TABLE DRIVEN TESTS
func TestAreaWithTable(t *testing.T) {
	areaTests := []struct {
		name   string
		shape  Shape
		expect float64
	}{
		{name: "rectangle", shape: Rectangle{Width: 20, Height: 30}, expect: 50.0},
		{"circle", Circle{8}, 201.06192982974676},
		// Pretty easy to add and test a new shape e.g triangle
		{"triangle", Triangle{12, 6}, 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.expect {
				t.Errorf(
					"at %#v:\n - got: %g \n - expected: %g",
					tt.shape, got, tt.expect,
				)
			}
		})
	}
}
