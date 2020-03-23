package main

func reverseWords(s string) string {
	result := ""
	subWord := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {

			subWord = []byte{}
		} else {
			subWord = append(subWord, s[i])
		}
	}
	result += reverseSubWord(string(subWord))
	return result
}

func reverseSubWord(subWord string) string {
	result := []byte{}
	for i := len(subWord) - 1; i >= 0; i-- {
		result = append(result, subWord[i])
	}
	return string(result)
}
