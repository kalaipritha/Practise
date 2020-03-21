package main

import (
	"fmt"
	s "strings"
)

func median(arr []int) {
}

func main() {
	str := "Te qick brown fox jumps over te lazy dog "
	var c int
	for i := 'a'; i <= 'z'; i++ {
		c = s.Count(str, string(i))
		if c < 1 {
			fmt.Println("not a panagram bcoz ", string(i), "is missing")
			c = -1
			break
		}
	}
	if c != -1 {
		fmt.Println("is a pangram")
	}
}
