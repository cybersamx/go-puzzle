package main

import "fmt"

func isPalindrome(s string) bool {
	length := len(s)

	if length == 0 {
		return false
	}

	var i, j int

	mid := length / 2

	if length%2 == 0 {
		// Even
		i = mid
		j = mid - 1
	} else {
		// Odd
		i = mid + 1
		j = mid - 1
	}

	for {
		if i >= length {
			break
		}

		if s[i] != s[j] {
			return false
		}

		i++
		j--
	}

	return true
}

func main() {
	positives := []string{
		"racecar",
		"madam",
		"naan",
		"radar",
		"kayak",
		"123456654321",
		"098767890",
	}

	negatives := []string{
		"",
		"lunch",
		"word",
		"sam",
	}

	for _, word := range positives {
		fmt.Printf("word: %s, isPalindrome: %v\n", word, isPalindrome(word))
	}

	for _, word := range negatives {
		fmt.Printf("word: %s, isPalindrome: %v\n", word, isPalindrome(word))
	}
}
