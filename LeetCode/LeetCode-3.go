package main

import "strings"

//滑动窗口
func lengthOfLongestSubstring(s string) int {
	length := 0
	left := 0
	right := 0
	s1 := s[left:right]

	for ; right < len(s); right++ {
		if currentIndex := strings.IndexByte(s1, s[right]); currentIndex != -1 {
			left += currentIndex + 1
		}
		s1 = s[left:right]
		if len(s1) > length {
			length = len(s1)
		}
	}
	return length
}
