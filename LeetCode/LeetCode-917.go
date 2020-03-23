package main

func reverseOnlyLetters(S string) string {
	sb := []byte(S)
	i := 0
	j := len(sb) - 1

	for i <= j {
		if !isLetter(sb[i]) {
			i++
		} else if !isLetter(sb[j]) {
			j--
		} else if isLetter(sb[i]) && isLetter(sb[j]) {
			tmp := sb[i]
			sb[i] = sb[j]
			sb[j] = tmp
			i++
			j--
		}
	}

	return string(sb)
}

func isLetter(char byte) bool {
	if (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') {
		return true
	}
	return false
}
