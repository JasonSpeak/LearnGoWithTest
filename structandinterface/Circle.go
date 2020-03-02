package structandinterface

import "math"

//Circle define a circle
type Circle struct {
	Radius float64
}

//Area get the area of a circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
