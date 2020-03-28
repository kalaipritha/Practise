package main

import (
	"fmt"
	"math"
)

func max(i int, j int) int {
	if i > j {
		return i
	}

	return j
}

func LongestBalancedSubstring(str1 string) {
	var count_1, count_0, res int
	m := make(map[int]int)
	m[0] = -1
	for i, s1 := range str1 {
		if string(s1) == "1" {
			count_1++
		} else {
			count_0++
		}
		diff := int(math.Abs(float64(count_1 - count_0)))
		_, found := m[diff]
		if found {
			res = max(res, i-m[diff])
		} else {
			m[diff] = i
		}
	}

	fmt.Println("the length if the longest balanced substring is", res)
}

func main() {
	str1 := "0100110001"
	LongestBalancedSubstring(str1)
}
