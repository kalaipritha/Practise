package main

import "fmt"

func uncomommonelement(arr []int, arr1 []int) {
	n := len(arr)
	m := len(arr1)
	i := 0
	j := 0
	for i < n && j < m {
		if arr[i] < arr1[j] {
			fmt.Println(arr[i])
			i++
		} else if arr1[j] < arr[i] {
			fmt.Println(arr1[j])
			j++
		} else {
			i++
			j++
		}
	}
	for i < n {
		fmt.Println(arr[i])
		i++
	}
	for j < m {
		fmt.Println(arr1[j])
		j++
	}
}

func main() {
	arr := []int{10, 20, 30}
	arr1 := []int{40, 50}
	uncomommonelement(arr, arr1)
}
