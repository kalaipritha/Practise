package main

import "fmt"

func swap(a, b int) (int, int) {
	return b, a
}

func reverse(arr []int, start int, end int) []int {
	if start < end {
		arr[start], arr[end] = swap(arr[start], arr[end])
		reverse(arr, start+1, end-1)

	}
	return arr
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	arr1 := reverse(arr, 0, len(arr)-1)
	fmt.Println(arr1)
}
