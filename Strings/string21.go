package main

import (
	"fmt"
	s "strings"
)

var top int = -1
var stack [20]int

func push(value int) {
	top++
	stack[top] = value

}

func pop() int {
	if top == -1 {
		return -1
	} else {
		val := int(stack[top])
		top--
		return val
	}
}

func CostTobalance(str1 string) {
	if s.Count(str1, "(") == s.Count(str1, ")") {
		var count int
		for i := len(str1) - 1; i >= 0; i-- {
			if string(str1[i]) == "(" {
				push(i)
			} else {
				val := pop()
				if val != -1 {
					count = count + (val - i)
				}
			}
		}
		fmt.Println("the cost to blance the string is ", count)
	} else {
		fmt.Println("the cost to blance the string is -1")
	}
}

func main() {
	str1 := "))(("
	CostTobalance(str1)
}
