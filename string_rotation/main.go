package main

import "fmt"

func isRotated(s1 string, s2 string) bool {
	// s1 and s2 have to be equal length and not empty
	if len(s1) != len(s2) || s1 == "" {
		return false
	}

	// Split the string into 2 substrings, swap/concatenate the 2
	// substrings and then compare
	for i := 0; i < len(s1); i++ {
		// Split the string into 2 substrings
		a := s1[i:]
		b := s1[:i]

		// Form a concatenated string
		ba := a + b
		if ba == s1 {
			return true
		}
	}

	return false
}

func main() {
	positives := map[string]string{
		"rotation": "rotation",
		"otationr": "rotation",
		"tationro": "rotation",
		"ionrotat": "rotation",
		"tionrota": "rotation",
		"ationrot": "rotation",
	}

	negatives := map[string]string{
		"": "rotation",
		"assign": "rotation",
		"shift": "rotation",
	}

	for k, v := range positives {
		fmt.Printf("s1: %s, s2: %s, isRotated: %v\n", v, k, isRotated(v, k))
	}

	for k, v := range negatives {
		fmt.Printf("s1: %s, s2: %s, isRotated: %v\n", v, k, isRotated(v, k))
	}
}
