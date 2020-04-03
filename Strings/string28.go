package main

import "fmt"

func FirstRepeatedletter(str1 string) string {
	word_count := make(map[string]int)
	for _, s := range str1 {
		if word_count[string(s)] == 0 {
			word_count[string(s)]++
		} else {
			return string(s)
		}
	}
	return "no repitition"
}

func main() {
	str1 := "hello geeks"
	res := FirstRepeatedletter(str1)
	if res != "no repitition" {
		fmt.Println("the first repeated letter is ", res)
	} else {
		fmt.Println("NO REPITITION")
	}
}
