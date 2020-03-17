package main

import "fmt"

func swap(a, b int) (int, int) {
	return b, a
}

func checkforsortedandrotated(arr []int) string {
	n := len(arr)
	var l, r int
	for i := 0; i < n-1; i++ {
		if arr[i] > arr[i+1] {
			fmt.Println("here 1")
			l = (i + 1) % n
			r = i
			break
			fmt.Println("l= r= ", l, r)
		}
	}

	if l == 0 && r == 0 {
		return "no"

	}
	fmt.Println(l, r)
	for l != r {
		if arr[r] < arr[l] {
			return "no"
		} else {
			l = (l + 1) % n
			r = (n + r - 1) % n
		}
	}
	return "yes"
}

func main() {
	arr := []int{7, 9, 11, 12, 5}
	str := checkforsortedandrotated(arr)
	fmt.Println(str)
}
