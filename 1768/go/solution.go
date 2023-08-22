package Solution1768

import (
	"strings"
)

func solution(word1 string, word2 string) string {
	len_word1 := len(word1)
	len_word2 := len(word2)

	var output strings.Builder

	var shorter_len int
	var longer_word string

	if len_word1 < len_word2 {
		shorter_len = len_word1
		longer_word = word2
	} else {
		shorter_len = len_word2
		longer_word = word1
	}

	for i := 0; i < shorter_len; i++ {
		output.WriteByte(word1[i])
		output.WriteByte(word2[i])
	}
	output.WriteString(longer_word[shorter_len:])

	return output.String()
}
