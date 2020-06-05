package main

// import "fmt"

type elements struct {
	i int
	j int
}

func enqueue(queue []elements, element elements) []elements {
	queue = append(queue, element)
	return queue
}

func dequeue(queue []elements) ([]elements, elements) {
	element := queue[0]
	return queue[1:], element
}

// func rottenOranges(matrix [][]int, queue []elements) {
// 	var count int
// 	for i := 0; i < 3; i++ {
// 		for j := 0; j < 5; j++ {
// 			if matrix[i][j] == 2 {
// 				queue = enqueue(queue, elements{
// 					i: i,
// 					j: j,
// 				})
// 			}
// 		}
// 	}
// 	queue = enqueue(queue, elements{
// 		i: -1,
// 		j: -1,
// 	})

// 	for len(queue) > 0 {
// 		var ele elements
// 		queue, ele = dequeue(queue)
// 		if ele.i != -1 && ele.j != -1 {
// 			if ele.i < 2 {
// 				if matrix[ele.i+1][ele.j] == 1 {
// 					matrix[ele.i+1][ele.j] = 2
// 					queue = enqueue(queue, elements{
// 						i: ele.i + 1,
// 						j: ele.j,
// 					})
// 				}
// 			}
// 			if ele.j < 4 {
// 				if matrix[ele.i][ele.j+1] == 1 {
// 					matrix[ele.i][ele.j+1] = 2
// 					queue = enqueue(queue, elements{
// 						i: ele.i,
// 						j: ele.j + 1,
// 					})
// 				}
// 			}
// 			if ele.i > 0 {
// 				if matrix[ele.i-1][ele.j] == 1 {
// 					matrix[ele.i-1][ele.j] = 2
// 					queue = enqueue(queue, elements{
// 						i: ele.i - 1,
// 						j: ele.j,
// 					})
// 				}
// 			}
// 			if ele.j > 0 {
// 				if matrix[ele.i][ele.j-1] == 1 {
// 					matrix[ele.i+1][ele.j-1] = 2
// 					queue = enqueue(queue, elements{
// 						i: ele.i,
// 						j: ele.j - 1,
// 					})
// 				}
// 			}
// 		} else {
// 			if len(queue) > 0 {
// 				queue = enqueue(queue, elements{
// 					i: -1,
// 					j: -1,
// 				})
// 				count++
// 			} else {
// 				fmt.Println("the required time is ", count)
// 			}
// 		}
// 	}
// }

// func main() {
// 	var queue []elements

// 	matrix := [][]int{{2, 1, 0, 2, 1},
// 		{1, 0, 1, 2, 1},
// 		{1, 0, 0, 2, 1}}

// 	rottenOranges(matrix, queue)

// }
