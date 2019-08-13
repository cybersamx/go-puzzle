package main

import (
	"bufio"
	"fmt"
	"os"
)

var lookupTable = []string{
	"academia",
	"academic",
	"account",
	"amazon",
	"amber",
	"ambassador",
	"apple",
}

func substring(s string, start int, length int) string {
	return s[start:length]
}

func matchingWords(str string) []string {
	var matches []string

	if str == "" {
		return matches
	}

	for _, word := range lookupTable {
		strSize := len(str)

		// Move to the next word if it's shorter than the user input
		if strSize > len(word) {
			continue
		}

		a := substring(str, 0, strSize)
		b := substring(word, 0, strSize)
		if a == b {
			matches = append(matches, word)
		}
	}

	return matches
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		str := scanner.Text()
		fmt.Println("You entered:", str)
		fmt.Println("Matching:")
		matches := matchingWords(str)
		for _, match := range matches {
			fmt.Println(match)
		}
	}
}
