package main

import (
	"fmt"
	"math"
)

func smallest(matrix [][]rune, i int, j int) int {
	min := math.MaxInt16
	if i <= 3 {
		if matrix[i+1][j] != 'O' && matrix[i+1][j] != 'W' {
			if min > int(matrix[i+1][j]) {
				min = int(matrix[i+1][j])
			}
		}
	}
	if i > 0 {
		if matrix[i-1][j] != 'O' && matrix[i-1][j] != 'W' {
			if min > int(matrix[i-1][j]) {
				min = int(matrix[i-1][j])
			}
		}

	}
	if j <= 3 {
		if matrix[i][j+1] != 'O' && matrix[i][j+1] != 'W' {
			if min > int(matrix[i][j+1]) {
				min = int(matrix[i][j+1])
			}
		}

	}
	if j > 0 {
		if matrix[i][j-1] != 'O' && matrix[i][j-1] != 'W' {
			if min > int(matrix[i][j-1]) {
				min = int(matrix[i][j-1])
			}
		}

	}
	return min + 1
}

func FindDistance(matrix [][]rune, queue []elements) {
	var mat [5][5]int

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if matrix[i][j] == 'G' {
				mat[i][j] = 1
				queue = enqueue(queue, elements{
					i: i,
					j: j,
				})
				matrix[i][j] = rune(0)
			}
		}
	}
	for len(queue) > 0 {
		var ele elements
		queue, ele = dequeue(queue)
		if matrix[ele.i][ele.j] == '0' {
			if ele.i <= 3 && matrix[ele.i+1][ele.j] != 'W' {
				matrix[ele.i+1][ele.j] = rune(1)
				mat[ele.i+1][ele.j] = 1
				queue = enqueue(queue, elements{
					i: ele.i + 1,
					j: ele.j,
				})
			}
			if ele.i > 0 && matrix[ele.i-1][ele.j] != 'W' {
				matrix[ele.i-1][ele.j] = rune(1)
				mat[ele.i-1][ele.j] = 1
				queue = enqueue(queue, elements{
					i: ele.i - 1,
					j: ele.j,
				})
			}
			if ele.j <= 3 && matrix[ele.i][ele.j+1] != 'W' {
				matrix[ele.i][ele.j+1] = rune(1)
				mat[ele.i][ele.j+1] = 1
				queue = enqueue(queue, elements{
					i: ele.i,
					j: ele.j + 1,
				})
			}
			if ele.j > 0 && matrix[ele.i][ele.j-1] != 'W' {
				matrix[ele.i][ele.j-1] = rune(1)
				mat[ele.i][ele.j-1] = 1
				queue = enqueue(queue, elements{
					i: ele.i,
					j: ele.j - 1,
				})
			}
		} else if matrix[ele.i][ele.j] == 'W' {
			continue
		} else {
			if ele.i <= 3 {
				if mat[ele.i+1][ele.j] != 1 && matrix[ele.i+1][ele.j] != 'W' {
					matrix[ele.i+1][ele.j] = rune(smallest(matrix, ele.i+1, ele.j))
					mat[ele.i+1][ele.j] = 1
					queue = enqueue(queue, elements{
						i: ele.i + 1,
						j: ele.j,
					})
				}

			}
			if ele.i > 0 {
				if mat[ele.i-1][ele.j] != 1 && matrix[ele.i-1][ele.j] != 'W' {
					matrix[ele.i-1][ele.j] = rune(smallest(matrix, ele.i-1, ele.j))
					mat[ele.i-1][ele.j] = 1
					queue = enqueue(queue, elements{
						i: ele.i - 1,
						j: ele.j,
					})
				}
			}
			if ele.j <= 3 {
				if mat[ele.i][ele.j+1] != 1 && matrix[ele.i][ele.j+1] != 'W' {
					matrix[ele.i][ele.j+1] = rune(smallest(matrix, ele.i, ele.j+1))
					mat[ele.i][ele.j+1] = 1
					queue = enqueue(queue, elements{
						i: ele.i,
						j: ele.j + 1,
					})
				}
			}
			if ele.j > 0 {
				if mat[ele.i][ele.j-1] != 1 && matrix[ele.i][ele.j-1] != 'W' {
					matrix[ele.i][ele.j-1] = rune(smallest(matrix, ele.i, ele.j-1))
					mat[ele.i][ele.j-1] = 1
					queue = enqueue(queue, elements{
						i: ele.i,
						j: ele.j - 1,
					})
				}

			}
		}
	}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if matrix[i][j] == 'W' {
				matrix[i][j] = rune(-1)
			}
		}
	}

	fmt.Println(matrix)

}

func main() {
	var queue []elements
	matrix := [][]rune{
		{'O', 'O', 'O', 'O', 'G'},
		{'O', 'W', 'W', 'O', 'O'},
		{'O', 'O', 'O', 'W', 'O'},
		{'G', 'W', 'W', 'W', 'G'},
		{'O', 'O', 'O', 'O', 'G'}}

	FindDistance(matrix, queue)
}
