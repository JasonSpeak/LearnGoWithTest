package main

import "fmt"

func reverseStr(s string, k int) string {
	sLength := len(s)
	res := ""
	for i := 0; i < sLength; {
		var ss string
		if i+2*k < sLength {
			ss = s[i : i+2*k]
		} else {
			ss = s[i:sLength]
		}
		if len(ss) >= k && len(ss) <= 2*k {
			res += reverseSubStr(ss[:k])
			res += ss[k:]
		} else {
			res += reverseSubStr(ss)
		}

		i += 2 * k
	}
	return res
}

func reverseSubStr(subStr string) string {
	result := []byte{}
	for i := len(subStr) - 1; i >= 0; i-- {
		result = append(result, subStr[i])
	}
	return string(result)
}

func main() {
	s := "abcdefg"
	res := reverseStr(s, 2)
	fmt.Println(res)
}
