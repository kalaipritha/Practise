package main

import (
	"fmt"
)

func combination(str string, inx int, curr_str string) {
	if inx == len(str) {
		fmt.Println(curr_str)
	} else {
		combination(str, inx+1, curr_str+string(str[inx]))
		combination(str, inx+1, curr_str)
	}
}

func main() {
	str1 := "aabc"
	combination(str1, 0, "")
}
