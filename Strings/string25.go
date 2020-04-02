package main

import (
	"fmt"
)

func LongestValidSunsequence(str1 string) {
	res := str1[0]
	for i, s1 := range str1 {
		if string(s1) == "A" {
			fmt.Println("the expression is ", res, "|", str1[i+1])
			res = res | str1[i+1]
			fmt.Println("the result is ", res)
		} else if string(s1) == "B" {
			fmt.Println("the expression is ", res, "&", str1[i+1])
			res = res & str1[i+1]
			fmt.Println("the result is ", res)
		} else if string(s1) == "C" {
			fmt.Println("the expression is ", res, "^", str1[i+1])
			res = res ^ str1[i+1]
			fmt.Println("the result is ", res)
		} else {
			continue
		}
	}
	fmt.Println(string(res))
}

func main() {
	str1 := "1A0B1"
	LongestValidSunsequence(str1)
}
