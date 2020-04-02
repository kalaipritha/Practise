package main

import (
	"fmt"
	s "strings"
)

func LongestValidSunsequence(str1 string) {
	if s.Count(str1, "[") == s.Count(str1, "]") {
		var count_left, count_right, swap, imbalance int
		for _, s1 := range str1 {
			if string(s1) == "[" {
				count_left++

				if imbalance > 0 {
					swap = swap + imbalance
					imbalance--
				}
			} else {
				count_right++

				imbalance = count_right - count_left
			}
		}
		fmt.Println("the number of required is ", swap)
	} else {
		fmt.Println("the number of required is -1")

	}
}

func main() {
	str1 := "[]][]["
	LongestValidSunsequence(str1)
}
