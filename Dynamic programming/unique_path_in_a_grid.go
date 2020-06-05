package main

// import "fmt"

// const (
// 	NO_OF_ROWS    = 3
// 	NO_OF_COLUMNS = 3
// )

// func unique_path(matrix [][]int) {
// 	var mat [3][3]int

// 	for j := 0; j < NO_OF_COLUMNS; j++ {
// 		if matrix[0][j] != 1 {
// 			mat[0][j] = 1
// 		} else {
// 			break
// 		}
// 	}

// 	for i := 0; i < NO_OF_ROWS; i++ {
// 		if matrix[i][0] != 1 {
// 			mat[i][0] = 1
// 		} else {
// 			break
// 		}
// 	}

// 	for i := 1; i < 3; i++ {
// 		for j := 1; j < 3; j++ {
// 			if matrix[i][j] != 1 {
// 				mat[i][j] = mat[i-1][j] + mat[i][j-1]
// 			}
// 		}
// 	}

// 	fmt.Println("the nnumber of unique paths are", mat[NO_OF_ROWS-1][NO_OF_COLUMNS-1])
// }

// func main() {
// 	matrix := [][]int{{0, 1, 0}, {0, 1, 0}, {0, 0, 0}}
// 	unique_path(matrix)

// }
