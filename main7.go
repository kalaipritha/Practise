package main

import "fmt"

func checkfornumberofjumps(arr []int) int {
	n := len(arr)
	arr2 := make([]int, n)
	arr3 := make([]int, n)
	arr3[0] = -1
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if arr[j] >= (i - j) {
				temp := arr2[j] + 1
				if temp < arr2[i] || arr2[i] == 0 {
					arr2[i] = temp
					arr3[i] = j
				}
			}
		}
	}

	return arr2[n-1]
}

func main() {
	arr := []int{1, 3, 6, 1, 0, 9}
	str := checkfornumberofjumps(arr)
	fmt.Println("the minimum jump required is", str)
}
