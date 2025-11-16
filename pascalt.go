package main

import (
	"fmt"
)

// computePascal returns a 2D slice containing n rows of Pascal's triangle
func computePascal(n int) [][]int {
	triangle := make([][]int, n)

	for i := 0; i < n; i++ {
		triangle[i] = make([]int, i+1)
		triangle[i][0] = 1
		triangle[i][i] = 1

		for j := 1; j < i; j++ {
			triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
		}
	}
	return triangle
}

func main() {
	n := 10 // Number of rows — change this value as needed

	triangle := computePascal(n)

	// Print Pascal triangle centered
	for i := 0; i < n; i++ {
		// spacing
		for s := 0; s < n-i; s++ {
			fmt.Print(" ")
		}

		// numbers
		for j := 0; j <= i; j++ {
			fmt.Printf("%4d", triangle[i][j])
		}
		fmt.Println()
	}
}
