package main

import (
	"fmt"
	s "strings"
)

const (
	UintSize = 32 << (^uint(0) >> 32 & 1)
	MaxInt   = 1<<(UintSize-1) - 1
)

func FindSmallestWindowOfSubstring(str string, pattern string) {
	pattern_hash := make(map[rune]int)
	str_hash := make(map[rune]int)
	pattern_length := len(pattern)
	var i int
	min_window_length := MaxInt
	for _, char := range pattern {
		if pattern_hash[char] == 0 {
			pattern_hash[char] = s.Count(pattern, string(char))
		}
	}

	for j, char := range str {
		str_hash[char]++
		if str_hash[char] <= pattern_hash[char] {
			pattern_length--
		}

		if pattern_length == 0 {
			for (pattern_hash[rune(str[i])] == 0 || str_hash[rune(str[i])] >= pattern_hash[rune(str[i])]) && pattern_length == 0 {
				if str_hash[rune(str[i])] > pattern_hash[rune(str[i])] && pattern_hash[rune(str[i])] != 0 {
					str_hash[rune(str[i])]--
					i++
				} else if str_hash[rune(str[i])] == pattern_hash[rune(str[i])] {
					str_hash[rune(str[i])]--
					pattern_length++
					i++
				} else {
					i++
				}
			}
			length := j - i + 1
			if min_window_length > length {
				min_window_length = length
			}
		}

	}

	fmt.Println("The minimun window length is ", min_window_length, "and the starting index is ", i-1)
}

func main() {
	str1 := "this is a test sring"
	fmt.Println("the given string is ", str1)
	fmt.Println("the pattern is tist")
	FindSmallestWindowOfSubstring(str1, "tist")
}
