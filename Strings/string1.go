package main

import (
	"fmt"
	s "strings"
)

func main() {
	str := "The qick brown fox jumps over lazy dog "
	str = s.ToLower(str)
	var c int
	for i := 'a'; i <= 'z'; i++ {
		c = s.Count(str, string(i))
		if c < 1 {
			fmt.Println("not a panagram ")
			c = -1
			break
		}
	}
	if c != -1 {
		fmt.Println("is a pangram")
	}
}
