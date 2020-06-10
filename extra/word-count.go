package main

import (
	"fmt"
	"strings"
)

func main() {
	text := `
    Needles and pins
    Needles and pins
    Sew me a sail
    To catch me the wind
    `
	words := strings.Fields(strings.ToLower(text))
	wordsCount := map[string]int{}

	for _, word := range words {
		wordsCount[word]++
	}

	fmt.Println(wordsCount)
}
