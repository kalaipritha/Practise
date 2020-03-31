package main

func push(value rune) {
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
