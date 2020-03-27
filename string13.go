package main

import (
	"fmt"
	s "strings"
)

func FindAnagram(str1 string, str2 string) {
	for _, s1 := range str2 {
		if s.Count(str1, string(s1)) != s.Count(str2, string(s1)) {
			fmt.Println(str2, " is not an anagram of ", str1)
			return
		}
	}

	fmt.Println(str2, " is an anagram of ", str1)
}

func main() {
	str1 := "geeksforgeeks"
	str2 := "forgeeksgeeks"
	FindAnagram(str1, str2)
}
