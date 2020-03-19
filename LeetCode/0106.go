package main

import (
	"fmt"
	"strconv"
)

func compressString(S string) string {
	SLen := len(S)
	if SLen < 2 {
		return S
	}
	currentChar := S[0]
	currentSum := 0
	result := ""
	for i := 0; i < SLen; i++ {
		if S[i] == currentChar {
			currentSum++
		} else {
			result += string(currentChar) + strconv.Itoa(currentSum)
			currentChar = S[i]
			currentSum = 1
		}
	}

	result += string(currentChar) + strconv.Itoa(currentSum)

	resultLen := len(result)
	if resultLen >= SLen {
		return S
	} else {
		return result
	}
}

func main() {
	S := "aabcccccaaa"
	res := compressString(S)
	fmt.Println(res)
}
