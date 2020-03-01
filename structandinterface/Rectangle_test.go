package structandinterface

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("test perimeter", func(t *testing.T) {
		got := Perimeter(Rectangle{10.00, 10.00})
		want := 40.00

		if got != want {
			t.Errorf("Got %.2f but want %.2f", got, want)
		}
	})
}

func TestArea(t *testing.T) {
	got := Area(Rectangle{10.00, 10.00})
	want := 100.00

	if got != want {
		t.Errorf("got %.2f but want %.2f", got, want)
	}
}
