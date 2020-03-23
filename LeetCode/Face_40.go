package main

import (
	"fmt"
	"sort"
)

func getLeastNumbers(arr []int, k int) []int {
	sort.Ints(arr)
	return arr[:k]
}

func main() {
	array := []int{3, 2, 1}
	res := getLeastNumbers(array, 2)
	fmt.Printf("%v", res)
}
