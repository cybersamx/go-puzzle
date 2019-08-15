package main

import "fmt"

// Given an integer n, convert it to binary, and then
// return the count of the longest sequence of zeros

func binarygap(n int) int {
	// Convert to binary as a slice of runes.
	binStr := fmt.Sprintf("%b", n)

	// Now count the number of zeros/sequences
	count := 0
	maxCount := 0

	for i, ch := range binStr {
		if ch == '0' {
			if i > 0 && []rune(binStr)[i-1] != ch {
				count = 0
			}

			count++
		} else if ch == '1' {
			if i > 0 && []rune(binStr)[i-1] != ch {
				if count > maxCount {
					maxCount = count
				}
			}
		}
	}

	if count > maxCount {
		maxCount = count
	}

	return maxCount
}

func main() {
	numbers :=  []int{ 0, 1, 15, 17, 84, 256, 33473, 7845500 }

	for _, n := range numbers {
		count := binarygap(n)
		fmt.Printf("Dec: %d, Bin: %b, Count: %d\n", n, n, count)
	}
}
