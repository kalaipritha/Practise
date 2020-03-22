package main

import "fmt"

func main() {
	str1 := "AXY"
	str2 := "YADXCP"
	i, j := 0, 0
	for i < len(str2) && j < len(str1) {
		if str1[j] == str2[i] {
			i++
			j++
		} else {
			i++
		}
	}

	if i < len(str2) {
		fmt.Println("false")
	} else {
		fmt.Println("true")
	}
}
