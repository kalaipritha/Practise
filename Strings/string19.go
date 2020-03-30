package main

import "fmt"

func EleminatingVowelndex(str1 string) {
	i := 0
	j := len(str1) - 1
	for i < len(str1) {
		if string(str1[i]) != "a" && string(str1[i]) != "e" && string(str1[i]) != "i" && string(str1[i]) != "o" && string(str1[i]) != "u" {
			fmt.Print(string(str1[j]))
		}
		j--
		i++
	}
}

func main() {
	str1 := "geeksforgeeks"
	EleminatingVowelndex(str1)
}
