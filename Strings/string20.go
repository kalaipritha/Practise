package main

import "fmt"

var top int = -1
var stack [2000]rune

func LongestValidSubstring(str1 string) {
	var count int

	for _, s := range str1 {
		if string(s) == "(" {
			push(s)
		} else {
			val := pop()
			if val != -1 {
				count = count + 2
			}
		}
	}
	fmt.Println("the longest valid substrings are ", count)
}

func main() {
	str1 := "))((("
	LongestValidSubstring(str1)
}
