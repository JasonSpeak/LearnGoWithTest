package main

func isAlienSorted(words []string, order string) bool {
	dictMap := make(map[byte]int)
	for i := 0; i < 26; i++ {
		dictMap[order[i]] = i
	}

	for i := 1; i < len(words); i++ {
		if !compareWords(words[i-1], words[i], dictMap) {
			return false
		}
	}

	return true

}

func compareWords(word1, word2 string, dictMap map[byte]int) bool {
	for i := 0; i < len(word1) && i < len(word2); i++ {
		if word1[i] != word2[i] {
			return dictMap[word1[i]] <= dictMap[word2[i]]
		}
	}
	return len(word1) <= len(word2)
}
