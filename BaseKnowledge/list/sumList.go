package list

func sum(array [5]int) int {
	sum := 0
	for _, num := range array {
		sum += num
	}
	return sum
}

//SumOfSlice returns sum of any size arary
func SumOfSlice(array []int) int {
	sum := 0
	for _, num := range array {
		sum += num
	}
	return sum
}

//SumAll returns Sums of every array
func SumAll(arrays ...[]int) (sums []int) {
	// sums = make([]int, len(arrays))

	// for i, array := range arrays {
	// 	sums[i] = SumOfSlice(array)
	// }
	for _, array := range arrays {
		sums = append(sums, SumOfSlice(array))
	}
	return sums
}

//SumAllTails returns all slices sum without the first num
func SumAllTails(arrays ...[]int) (sums []int) {
	for _, array := range arrays {
		if len(array) == 0 {
			sums = append(sums, 0)
		} else {
			tails := array[1:]
			sums = append(sums, SumOfSlice(tails))
		}
	}
	return
}
