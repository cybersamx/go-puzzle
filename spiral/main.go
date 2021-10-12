package main

import "fmt"

func newTestMatrices() [][][]int {
	var matrixEmpty = [][]int{{}}
	var matrix1 = [][]int{{1}}
	var matrix2 = [][]int{
		{1, 2},
		{4, 3},
	}

	var matrixRow = [][]int{
		{1, 2, 3, 4},
	}

	var matrixCol = [][]int{
		{1},
		{2},
		{3},
		{4},
	}

	var matrix4 = [][]int{
		{1, 2, 3, 4},
		{12, 13, 14, 5},
		{11, 16, 15, 6},
		{10, 9, 8, 7},
	}

	var matrix5 = [][]int{
		{1, 2, 3, 4, 5},
		{16, 17, 18, 19, 6},
		{15, 24, 25, 20, 7},
		{14, 23, 22, 21, 8},
		{13, 12, 11, 10, 9},
	}

	return [][][]int{
		matrixEmpty,
		matrix1,
		matrix2,
		matrixRow,
		matrixCol,
		matrix4,
		matrix5,
	}
}

func checkSquareMatri(matrix [][]int) bool {
	if len(matrix) == 0 {
		return false
	}

	width := len(matrix[0])
	for row := range matrix {
		if row > 0 && width != len(matrix[row]) {
			return false
		}

		width = len(matrix[row])
	}

	return width == len(matrix)
}

func spiralTraverse(matrix [][]int) {
	if !checkSquareMatri(matrix) {
		fmt.Println("Empty or not square matrix")
		return
	}

	width := len(matrix[0])
	height := len(matrix)

	col := 0
	row := 0

	for i := 0; i < width+height-1; i++ {
		n := i % 4

		if n == 0 {
			// Left to right
			col = i / 4
			row = i / 4

			for ; col < width-i/4; col++ {
				fmt.Printf("%d ", matrix[row][col])
			}
		} else if n == 1 {
			// Top to bottom
			col = width - i/4 - 1
			row = i/4 + 1
			for ; row < height-i/4; row++ {
				fmt.Printf("%d ", matrix[row][col])
			}
		} else if n == 2 {
			// Right to left
			col = width - i/4 - 2
			row = height - i/4 - 1
			for ; col >= i/4; col-- {
				fmt.Printf("%d ", matrix[row][col])
			}
		} else if n == 3 {
			// Bottom to top
			col = i / 4
			row = height - i/4 - 2
			for ; row > i/4; row-- {
				fmt.Printf("%d ", matrix[row][col])
			}
		}
	}

	fmt.Println()
}

func main() {
	for _, matrix := range newTestMatrices() {
		spiralTraverse(matrix)
	}
}
