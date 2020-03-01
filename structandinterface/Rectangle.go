package structandinterface

//Rectangle defines struct of rectangle
type Rectangle struct {
	width  float64
	height float64
}

//Perimeter returns the perimeter of a rectangle
func Perimeter(rect Rectangle) float64 {
	if rect.width <= 0.0 || rect.height <= 0.0 {
		return 0.0
	}
	return 2.0 * (rect.height + rect.width)
}

//Area returns the area of a rectangle
func Area(rect Rectangle) float64 {
	if rect.width <= 0.0 || rect.height <= 0.0 {
		return 0.0
	}
	return rect.height * rect.width
}
