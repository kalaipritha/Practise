package main

import "fmt"

func max(i int, j int) int {
	if i > j {
		return i
	}

	return j
}

func reverse(str string, strt int, end int) string {
	str2 := ""
	for i := strt; i >= end; i-- {
		str2 = str2 + string(str[i])
	}
	return str2
}

func StringLeftRotation(str1 string, d int) string {
	return reverse((reverse(str1, d-1, 0) + reverse(str1, len(str1)-1, d)), len(str1)-1, 0)
}

func StringRightRotation(str1 string, d int) string {
	return reverse((reverse(str1, len(str1)-1-d, 0) + reverse(str1, len(str1)-1, len(str1)-d)), len(str1)-1, 0)
}

func main() {
	str1 := "GeeksforGeeks"
	str2 := StringLeftRotation(str1, 2)
	fmt.Println("left rotation is", str2)
	str2 = StringRightRotation(str1, 2)
	fmt.Println("right rotation is", str2)
}
