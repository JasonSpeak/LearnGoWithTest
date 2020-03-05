package structandinterface

//Triangle define a triangle
type Triangle struct {
	Base   float64
	Height float64
}

//Area returns the area of a triangle
func (triangle Triangle) Area() float64 {
	return triangle.Base * triangle.Height / 2
}
