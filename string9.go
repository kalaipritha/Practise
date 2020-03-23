package main

import (
	"fmt"
)

func FindSubsequence(str string, indx int, cur_str string, k string, count int) int {
	if indx == len(str) {
		if cur_str == k {
			count++
		}
		return count
	} else {
		count = FindSubsequence(str, indx+1, cur_str+string(str[indx]), k, count)
		count = FindSubsequence(str, indx+1, cur_str, k, count)
	}
	return count

}

func main() {
	str1 := "GeeksforGeeks"
	count := FindSubsequence(str1, 0, "", "Gks", 0)
	fmt.Println(count)
}
