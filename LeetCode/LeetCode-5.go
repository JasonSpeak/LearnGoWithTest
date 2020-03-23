package main

func longestPalindrome(s string) string {
	if s == nil || len(s) < 1 {
		return ""
	}
	start := 0
	end := 0

	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)
		length := max(len1, len2)
		if length > end-start {
			start = i - (length-1)/2
			end = i + length/2
		}
	}

	return s[start:end]

}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func expandAroundCenter(s string, left, right int) int {
	L := left
	R := right
	for L >= 0 && R < len(s) && s[L] == s[R] {
		L--
		R++
	}
	return R - L - 1
}
