package main

import "fmt"

func FirstRepeatedWord(str1 string) string {
	word_count := make(map[string]int)
	str2 := ""

	for _, s := range str1 {
		if string(s) == " " {
			word_count[str2]++
			str2 = ""
		} else {
			str2 = str2 + string(s)
		}
	}

	word_count[str2]++
	str2 = ""

	for _, s := range str1 {
		if string(s) == " " {
			if word_count[str2] > 1 {
				return str2
			}
			str2 = ""
		} else {
			str2 = str2 + string(s)
		}
	}

	return "no repitition"
}

func main() {
	str1 := "Ravi had been saying that"
	res := FirstRepeatedWord(str1)
	if res != "no repitition" {
		fmt.Println("the first repeated word is ", res)
	} else {
		fmt.Println("NO REPITITION")
	}
}
