package main

import (
	"strings"
)

func cleanInput(text string) []string {
	trimmed := strings.TrimSpace(text)
	splitted := strings.Split(trimmed, " ")
	words := []string{}
	for _, word := range splitted {
		if len(word) > 0 {
			words = append(words, strings.ToLower(word))
		}
	}

	return words
}
