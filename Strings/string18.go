package main

import "fmt"

func ReverseWordInString(str1 string) string {
	strt := len(str1) - 1
	str2 := ""
	for i := strt; i >= 0; i-- {
		if string(str1[i]) == " " {
			i1 := i + 1
			for i1 <= strt {
				str2 = str2 + string(str1[i1])
				i1++
			}
			str2 = str2 + " "
			strt = i - 1
		}
	}
	i1 := 0
	for i1 <= strt {
		str2 = str2 + string(str1[i1])
		i1++
	}
	return str2
}

func main() {
	str1 := "GfG IS THE BEST"
	str2 := ReverseWordInString(str1)
	fmt.Println("the reverse of a words in string is ", str2)
}
