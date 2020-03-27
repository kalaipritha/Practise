package main

import "fmt"

func BuildHeap(arr []int, k int) []int {
	for i := k/2 - 1; i >= 0; i-- {
		arr = Heapify(arr, i, k)
	}
	return arr
}

func Heapify(arr []int, i int, k int) []int {
	if i >= k/2 {
		return arr
	}

	var smallest int
	l := 2*i + 1
	r := 2*i + 2

	if arr[l] < arr[i] {
		smallest = l
	} else {
		smallest = i
	}

	if arr[r] < arr[smallest] {
		smallest = r
	}

	if smallest != i {
		temp := arr[smallest]
		arr[smallest] = arr[i]
		arr[i] = temp
		arr = Heapify(arr, smallest, k)
	}

	return arr
}

func LargestKelememts(arr []int, k int) []int {
	arr = BuildHeap(arr, k)
	for i := k; i < len(arr); i++ {
		if arr[0] > arr[i] {
			continue
		} else {
			temp := arr[0]
			arr[0] = arr[i]
			arr[i] = temp
			arr = Heapify(arr, 0, k)
		}
	}
	return arr
}

func main() {
	arr := []int{11, 3, 2, 1, 15, 5, 4,
		45, 88, 96, 50, 3}
	k := 5
	arr1 := LargestKelememts(arr, k)
	for i := 0; i < k; i++ {
		fmt.Println(arr1[i])
	}
}
