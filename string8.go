package main

import (
	"fmt"
	"strconv"
)

func FindSubsequence(str string, indx int, cur_str string, k int, count int) int {
	if indx == len(str) {
		num, _ := strconv.Atoi(cur_str)
		if num%k == 0 {
			fmt.Println("the number is ", num)
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
	str1 := "676"
	count := FindSubsequence(str1, 0, "", 6, 0)
	fmt.Println(count - 1)
}
