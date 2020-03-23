package main

import "sort"

func minIncrementForUnique(A []int) int {
	if len(A) == 0 {
		return 0
	}
	sort.Ints(A)
	count := 0
	for i := 1; i < len(A); i++ {
		if A[i] <= A[i-1] {
			pre := A[i]
			A[i] = A[i-1] + 1
			count += A[i] - pre
		}
	}
	return count
}
