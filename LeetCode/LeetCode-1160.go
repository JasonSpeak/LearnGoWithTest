package main

import "fmt"

func countCharacters(words []string, chars string) int {
	count := 0
	for i := 0; i < len(words); i++ {
		tmpAlphabet := getAlphabet(chars)
		for j := 0; j < len(words[i]); j++ {
			if tmpAlphabet[words[i][j]] > 0 {
				tmpAlphabet[words[i][j]]--
				if j == len(words[i])-1 {
					count += len(words[i])
				}
			} else {
				break
			}
		}
	}
	return count
}

func getAlphabet(chars string) map[byte]int {
	alphabet := map[byte]int{}
	for i := 0; i < len(chars); i++ {
		alphabet[chars[i]]++
	}
	return alphabet
}

func main() {
	words := []string{"cat", "bt", "hat", "tree"}
	chars := "atach"

	res := countCharacters(words, chars)

	fmt.Println(res)
}
