package main

import (
	"fmt"
	s "strings"
)

func main() {
	str1 := "gggg"
	if s.Count(str1, string(str1[0])) == len(str1) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
