package main

import "fmt"

func canThreePartsEqualSum(A []int) bool {
	sum := 0
	for i := 0; i < len(A); i++ {
		sum += A[i]
	}
	if sum%3 != 0 {
		return false
	}

	partSum := sum / 3
	lastPartIndex := -1
	blockCount := 0
	tmpSum := 0
	for i := 0; i < len(A); i++ {
		tmpSum += A[i]
		if tmpSum == partSum {
			lastPartIndex = i
			blockCount++
			tmpSum = 0
		}

		if blockCount == 2 {
			break
		}
	}

	fmt.Println("blocks is %d and lastIndex is %d", blockCount, lastPartIndex)
	if blockCount == 2 && lastPartIndex < len(A)-1 {
		tmp := 0
		for i := lastPartIndex + 1; i < len(A); i++ {
			tmp += A[i]
		}
		if tmp == partSum {
			return true
		}
	}

	return false
}

func main() {
	test := []int{10, -7, 13, -14, 30, 16, -3, -16, -27, 27, 19}
	result := canThreePartsEqualSum(test)
	fmt.Println(result)
}
