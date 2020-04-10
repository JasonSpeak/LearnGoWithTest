package main

import "fmt"

func reverseWords(s string) string {
	length := len(s)
	words := []string{}
	wordBytes := []byte{}
	result := ""
	for i := 0; i < length; i++ {
		if s[i] == ' ' && len(wordBytes) != 0 {
			words = append(words, string(wordBytes))
			wordBytes = []byte{}
		} else {
			wordBytes = append(wordBytes, s[i])
		}
	}
	words = append(words, string(wordBytes))
	wordsLength := len(words)
	for i := wordsLength - 1; i > 0; i-- {
		if len(words[i]) != 0 {
			result = result + words[i] + " "
		}
	}
	if words[0] != " " {
		result += words[0]
	}
	return result
}

func main() {
	fmt.Println(reverseWords("the sky is blue"))
}
