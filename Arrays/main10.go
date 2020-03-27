package main

import "fmt"

func MergeTwoSortedArray(arr []int, arr1 []int) {
	n := len(arr)
	m := len(arr1)
	i := n - 1
	for j := n - 1; j >= 0; j-- {
		if arr[j] != 0 {
			temp := arr[i]
			arr[i] = arr[j]
			arr[j] = temp
			i--
		}
	}

	j := i + 1
	k := 0
	l := 0

	for j < n && k < m {
		if arr[j] < arr1[k] {
			arr[l] = arr[j]
			j++
			l++
		} else if arr1[k] < arr[j] {
			arr[l] = arr1[k]
			k++
			l++
		} else {
			arr[l] = arr1[k]
			j++
			k++
			l++
		}
	}

	for j < n {
		arr[l] = arr[j]
		j++
		l++
	}

	for k < m {
		arr[l] = arr1[k]
		k++
		l++
	}

	fmt.Println("the merged array is ", arr)
}

func main() {
	arr := []int{0, 0, 0, 0, 0, 0, 1}
	arr1 := []int{5, 8, 12, 14}
	MergeTwoSortedArray(arr, arr1)
}
