package Question_by_day

import "strings"

func isPrefixOfWord(sentence string, searchWord string) int {
	words := strings.Split(sentence, " ")
	for idx, word := range words {
		if len(word) < len(searchWord) {
			continue
		}
		flag := true
		for i := 0; i < len([]byte(searchWord)); i++ {
			if word[i] != searchWord[i] {
				flag = false
				break
			}
		}
		if flag {
			return idx + 1
		}
	}
	return -1
}
