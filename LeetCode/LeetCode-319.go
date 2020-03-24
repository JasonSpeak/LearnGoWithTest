package main

func bulbSwitch(n int) int {
	if n == 1 {
		return 1
	}

	result := 1
	for true {
		if result*result > n {
			break
		}
		result++
	}
	return result - 1
}
