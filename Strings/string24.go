package main

import (
	"fmt"
	s "strings"
)

func LongestValidSunsequence(str1 string) {
	var max, curr_max int
	if s.Count(str1, "(") == s.Count(str1, ")") {
		for _, s1 := range str1 {
			if string(s1) == "(" {
				curr_max++
				if curr_max > max {
					max = curr_max
				}
			} else if string(s1) == ")" {
				curr_max--
			} else {
				continue
			}
		}
		fmt.Println("the maximum depth is ", max)
	} else {
		fmt.Println("-1")
	}
}

func main() {
	str1 := "( ((X)) (((Y))) )"
	LongestValidSunsequence(str1)
}
