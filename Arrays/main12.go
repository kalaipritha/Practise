package main

import (
	"fmt"
	"math"
)

func max(a float64, b float64) float64 {
	if a > b {
		return a
	} else {
		return b
	}
}

func search(arr []int, k int, x int) int {
	i := 0
	for i < len(arr) {
		if arr[i] == x {
			return i
		} else {
			i = i + int(max(1.0, math.Abs(float64(arr[i]-x))/float64(k)))
		}
	}
	return -1
}

func main() {
	arr := []int{20, 40, 50, 70, 70, 60}
	k := 20
	x := 60
	index := search(arr, k, x)
	fmt.Println(index)
}
