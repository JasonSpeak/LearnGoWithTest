package main

func canMeasureWater(x int, y int, z int) bool {
	if z == 0 {
		return true
	}

	if z > x+y {
		return false
	}

	for y != 0 {
		x, y = y, x%y
	}

	if x != 0 {
		if z%x == 0 {
			return true
		}
	}
	return false
}
