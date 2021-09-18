package str

import (
	"strings"
)

func AddQuotesBetweenWords(words []string, char string) (sentences string) {
	for _, word := range words {
		sentences += `'` + word + `'` + char
	}
	// remove char in the last word
	sentences = strings.TrimRight(sentences, char)
	return sentences
}

func AddCharBetweenWords(words []string, char string) (sentences string) {
	for _, word := range words {
		sentences += word + char
	}
	// remove char in the last word
	sentences = strings.TrimRight(sentences, char)
	return sentences
}

// add space between words
func Spacing(words ...string) (sentences string) {
	sentences = AddCharBetweenWords(words, " ")
	return sentences
}
