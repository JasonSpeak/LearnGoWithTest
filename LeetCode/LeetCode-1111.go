package main

func maxDepthAfterSplit(seq string) []int {
	length := len(seq)
	res := make([]int, length)

	for i := 0; i < length; i++ {
		if seq[i] == '(' {
			res[i] = i & 1
		} else {
			res[i] = (i + 1) & 1
		}
	}
	return res
}
