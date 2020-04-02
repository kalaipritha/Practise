package main

import "fmt"

var top int = -1
var stack [20]rune

func push(value rune) {
	top++
	stack[top] = value

}

func pop() rune {
	if top == -1 {
		return rune(-1)
	} else {
		val := stack[top]
		top--
		return val
	}
}

func CheckForBalancesdString(str1 string) {
	flag := true
	for _, s1 := range str1 {
		if string(s1) == "(" || string(s1) == "[" || string(s1) == "{" {
			push(s1)
		} else {
			val := pop()
			if (string(s1) == ")" && string(val) != "(") || (string(s1) == "]" && string(val) != "[") || (string(s1) == "}" && string(val) != "{") {
				fmt.Println("unbalanced paranthesis")
				flag = false
				break
			}
		}
	}
	if flag == true {
		fmt.Println("balanced paranthesis")
	}

}

func main() {
	str1 := "()]"
	CheckForBalancesdString(str1)
}
