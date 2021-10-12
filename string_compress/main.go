package main

import (
	"fmt"
	"strings"
)

func compress(str string) string {
	var sb strings.Builder

	runes := []rune(str)
	count := 1

	for i, ch := range runes {
		if i < len(runes)-1 {
			// Look 1 character ahead (but not the last one)
			c := runes[i+1]
			if c != ch {
				// Start of repeated char, reset count and append current char ch
				sb.WriteRune(ch)
				sb.WriteString(fmt.Sprintf("%d", count))
				count = 1
			} else {
				// Don't append repeated character
				count++
			}
		} else {
			// At last character
			sb.WriteRune(ch)
			sb.WriteString(fmt.Sprintf("%d", count))
		}
	}

	if sb.Len() >= len(str) {
		return str
	}

	return sb.String()
}

func main() {
	strings := []string{
		"a",
		"aa",
		"aaaa",
		"aaaabbbccd",
		"abcdddddd",
		"abcc",
	}

	for _, s := range strings {
		fmt.Println(compress(s))
	}
}
