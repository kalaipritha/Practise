package main

import (
	"fmt"
)

func max(i int, j int) int {
	if i > j {
		return i
	}

	return j
}

func SubstringContainingAllVowel(str1 string) {
	varmap := make(map[rune]int)
	str := ""
	for _, s1 := range str1 {
		if string(s1) == "a" || string(s1) == "e" || string(s1) == "i" || string(s1) == "o" || string(s1) == "u" {
			varmap[s1]++
			str = str + string(s1)
			if varmap['a'] >= 1 && varmap['u'] >= 1 && varmap['e'] >= 1 && varmap['i'] >= 1 && varmap['o'] >= 1 {
				fmt.Println(str)
			}

		} else {
			varmap = make(map[rune]int)
			str = ""
		}
	}
}

func main() {
	str1 := "aeoibsddaeiofuddudb"
	SubstringContainingAllVowel(str1)
}
