package main

import (
	"fmt"
	s "strings"
)

func check(str1 string) bool {
	count := s.Count(str1, "0")
	var flag bool
	if count > 0 {
		flag = true
	} else {
		flag = false
	}
	return flag && string(str1[0]) != "0"
}

func main() {
	str1 := "77609"
	if check(str1) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
