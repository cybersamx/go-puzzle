package main

import (
	"fmt"
)

func main() {
	str := "0123456789"

	// Slicing a slice.
	fmt.Println("str[:7] =>", str[:7])
	fmt.Println("str[3:] =>", str[3:])
	fmt.Println("str[3:5] =>", str[3:5])

	// String to bytes and vice versa.
	sliceByte := []byte(str)
	fmt.Println("byte slice =>", sliceByte)

	sliceRune := []rune(str)
	fmt.Println("rune slice =>", sliceRune)

	fmt.Println("back to string", string(sliceByte))
	fmt.Println("back to string", string(sliceRune))

	// Appending a slice.
	sliceByte = append(sliceByte, 'a')          // with 1 item
	sliceByte = append(sliceByte, []byte{'b', 'c'}...)  // with another slice of the same type
	fmt.Println("appended slice (string)", string(sliceByte))

	// Reference a character in string.
	fmt.Println("string at 4 =>", str[4])

	// Loops.
	for i, v := range str {
		fmt.Printf("%d-%c ", i, v)
	}
	fmt.Println()

	for i := 0; i < len(str); i++ {
		fmt.Printf("%d-%c ", i, str[i])
	}
	fmt.Println()
}
