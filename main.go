package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	words := readInput()
	words = filterEmptyWords(words)
	fmt.Println(len(words))
}

func readInput() []string {
	args := os.Args[1:]
	words := []string{}
	if len(args) != 0 {
		words = strings.Split(args[0], " ")
	}

	return words
}

func filterEmptyWords(words []string) []string {
	filterWords := []string{}
	for _, word := range words {
		if len(word) != 0 {
			filterWords = append(filterWords, word)
		}
	}

	return filterWords
}
