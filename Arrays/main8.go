package main

import "fmt"

func removeDuplicate(arr []int) {
	n := len(arr)
	j := 0
	for i := 0; i < n-1; i++ {
		if arr[i] != arr[i+1] {
			arr[j] = arr[i]
			j++
		}
	}
	arr[j] = arr[n-1]
	for k := 0; k <= j; k++ {
		fmt.Println(arr[k])
	}
}

func main() {
	arr := []int{2, 2, 2, 2, 2}
	removeDuplicate(arr)
}
