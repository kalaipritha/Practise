package main

import "fmt"

const (
	m = 3
	n = 3
)

func PrintAllPaths(matrix [][]int, arr *[5]int, i int, j int, idx int) {
	(*arr)[idx] = matrix[i][j]
	if j > 0 {
		PrintAllPaths(matrix, arr, i, j-1, idx+1)
	}

	if i > 0 {
		PrintAllPaths(matrix, arr, i-1, j, idx+1)
	}

	if i == 0 && j == 0 {
		fmt.Print((*arr))
		fmt.Println("")
	}
}

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	var arr [m + n - 1]int

	PrintAllPaths(matrix, &arr, 2, 2, 0)
}
