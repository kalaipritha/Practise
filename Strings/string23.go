package main

import "fmt"

func LongestValidSunsequence(str1 string) {
	var iob, icb int
	for _, s1 := range str1 {
		if string(s1) == "(" {
			iob++
		} else {
			if iob == 0 {
				icb++
			} else {
				iob--
			}
		}
	}
	fmt.Println(len(str1) - (iob + icb))
}

func main() {
	str1 := "()())"
	LongestValidSunsequence(str1)
}
