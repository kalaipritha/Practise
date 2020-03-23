package main

import (
	"fmt"
	s "strings"
)

func CheckPalindrome(str string) bool {
	i := 0
	j := len(str) - 1
	for i < j {
		if str[i] != str[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func FindRepeatedSubsequence(str string) bool {
	var s2 string
	for _, s1 := range str {
		if s.Count(str, string(s1)) > 2 {
			return true
		} else if s.Count(str, string(s1)) > 1 {
			s2 = s2 + string(s1)
		}
	}

	if CheckPalindrome(s2) {
		if len(s2)%2 == 1 && s2[len(s2)/2] == s2[(len(s2)/2)-1] && s2[len(s2)/2] == s2[(len(s2)/2)+1] {
			return true
		} else {
			return false
		}
	} else {
		return true
	}
}

func main() {
	str1 := "ABCD"
	count := FindRepeatedSubsequence(str1)
	fmt.Println(count)
}
