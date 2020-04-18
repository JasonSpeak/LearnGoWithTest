package main

func canJump(nums []int) bool {
	k := 0
	length := len(nums)
	for i := 0; i < length; i++ {
		if i > k {
			return false
		}
		k = max(k, i+nums[i])
	}
	return true
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
