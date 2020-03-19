package main

func lengthOfLIS(nums []int) int {
	len := len(nums)
	if len < 2 {
		return len
	}

	dp := []int{}
	for i := 0; i < len; i++ {
		dp = append(dp, 1)
	}

	for i := 1; i < len; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = Max(dp[i], dp[j]+1)
			}
		}
	}

	res := 0
	for i := 0; i < len; i++ {
		res = Max(res, dp[i])
	}

	return res
}

func Max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
