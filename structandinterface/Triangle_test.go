package structandinterface

import "testing"

func TestTriangle_Area(t *testing.T) {
	tests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "REctangle", shape: Rectangle{12, 6}, want: 72.0},
		{"Circle", Circle{10}, 314.1592653589793},
		{"Triangle", Triangle{12, 6}, 36.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.shape.Area(); got != tt.want {
				t.Errorf("Triangle.Area() = %v, want %v", got, tt.want)
			}
		})
	}
}
