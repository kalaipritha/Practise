package main

import (
	"fmt"
	s "strings"
)

func main() {
	str1 := "klxml"
	str2 := "kxml"
	for _, s1 := range str1 {
		if s.Count(str2, string(s1)) != s.Count(str1, string(s1)) {
			fmt.Println(string(s1))
			break
		}
	}
}
