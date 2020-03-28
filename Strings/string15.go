package main

import (
	"fmt"
)

func max(i int, j int) int {
	if i > j {
		return i
	}

	return j
}

func LongestVowelSubstring(str1 string) {
	var count, max_count int
	for _, s1 := range str1 {
		if string(s1) == "a" || string(s1) == "e" || string(s1) == "i" || string(s1) == "o" || string(s1) == "u" {
			count = count + 1
		} else {
			max_count = max(max_count, count)
			count = 0
		}
	}
	fmt.Println("the length of the longest vowel substring is", max_count)
}

func main() {
	str1 := "theeaere"
	LongestVowelSubstring(str1)
}
