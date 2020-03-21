package main

import (
	"fmt"
	s "strings"
)

func main() {
	str1 := "kalai pritha r"
	arr := s.SplitAfter(str1, " ")
	for _, s1 := range arr {
		fmt.Println(s.ToUpper(string(s1[0])))
	}
}
