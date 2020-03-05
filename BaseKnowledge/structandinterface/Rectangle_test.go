package structandinterface

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	t.Run("test perimeter", func(t *testing.T) {
		rectangle := Rectangle{10.00, 10.00}
		got := rectangle.Perimeter()
		want := 40.00

		if got != want {
			t.Errorf("Got %.2f but want %.2f", got, want)
		}
	})

}

func TestArea(t *testing.T) {
	rectangle := Rectangle{10.00, 10.00}
	got := rectangle.Area()
	want := 100.00

	if got != want {
		t.Errorf("got %.2f but want %.2f", got, want)
	}
}

func TestCircles(t *testing.T) {
	circle := Circle{10}
	got := circle.Area()
	want := 314.16

	if got != want {
		t.Errorf("got %.2f but want %.2f", got, want)
	}
}

func TestRectangle_Perimeter(t *testing.T) {
	tests := []struct {
		name string
		rect Rectangle
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.rect.Perimeter(); got != tt.want {
				t.Errorf("Rectangle.Perimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAreaWithInterface(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %.2f but want %.2f", got, want)
		}
	}

	t.Run("Rectangles", func(t *testing.T) {
		rectangle := Rectangle{10.00, 10.00}
		checkArea(t, rectangle, 100.00)
	})

	t.Run("Circles", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})
}

