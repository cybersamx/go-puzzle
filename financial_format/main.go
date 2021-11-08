package main

import (
	"fmt"
	"math"
)

const digits = "0123456789"


// Solve this problem mathematically.
//
// itoa converts the int parameter and returns the text representation.
// Given n = 1954. The math is:
// 1954 = 1000 => i=3 => 1954             => 1954/10^3 => floor(1.954) => 1.0 => 1
//         900 => i=2 => 1954-1000        =>  954/10^2 => floor(9.54)  => 9.0 => 9
//          50 => i=1 => 1954-1000-900    =>   54/10^1 => floor(5.4)   => 5.0 => 5
//           4 => i=0 => 1954-1000-900-50 =>    4/10^0 => floor(4.0)   => 4.0 => 4
func itoa(num int) string {
	if num == 0 {
		return "0"
	}

	var chars []byte

	numf := math.Abs(float64(num))
	prev := 0.0
	n := int(math.Log10(numf))

	if num < 0 {
		chars = append(chars, '-')
	}

	for i := n; i >= 0; i-- {
		numf -= prev
		m := math.Floor(numf / math.Pow10(i))
		prev = m * math.Pow10(i)

		if i%3 == 2 && i < n {
			chars = append(chars, ',')
		}

		chars = append(chars, digits[int(m)])
	}

	return string(chars)
}

// Solve this problem by string manipulation.
//
// itoastr basically convert the integer to a string and then insert the comma conditionally.
func itoastr(num int) string {
	// Convert to string.
	nums := fmt.Sprintf("%d", num)

	if num < 0 {
		nums = fmt.Sprintf("%d", num * -1)
	}

	count := len(nums)

	var chars []byte

	// We can have 3 states:
	// 123 (no comma) => firstComma = 0
	// 1,234 (comma at index 1) => firstComma = 1
	// 12,345 (comma at index 2) => firstComma = 2
	firstComma := count % 3

	for i := 0; i < count; i++ {
		chars = append(chars, nums[i])

		if i >= count - 1 {
			continue
		}

		if i <= 3 && firstComma >= 0 && firstComma == i + 1 {
			chars = append(chars, ',')
		} else if i >= 2 && (i + 1 - firstComma) % 3 == 0 {
			chars = append(chars, ',')
		}
	}

	if num < 0 {
		chars = append([]byte{'-'}, chars...)
	}

	return string(chars)
}

func main() {
	inputs := []int{0, 1, 123, 1234, 12345, 123456, 1234567, -1, -123, -1234}
	for _, input := range inputs {
		fmt.Println(itoa(input))
		fmt.Println(itoastr(input))
	}
}
