package structandinterface

import "testing"

func TestCircle_Area(t *testing.T) {
	tests := []struct {
		name string
		c    Circle
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Area(); got != tt.want {
				t.Errorf("Circle.Area() = %v, want %v", got, tt.want)
			}
		})
	}
}
