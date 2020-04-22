package main

import "fmt"

func (p *node) Reverse(h *list, n int) *list {
	current := h
	var next, prev *list
	count := 0
	for current != nil && count < n {
		next = current.next
		current.next = prev
		prev = current
		current = next
		count++
	}

	if next != nil {
		h.next = p.Reverse(next, n)
	}

	return prev
}

func main() {
	list := createList()
	arr := []int{3, 456, 7, 20, 4}
	for i := 0; i < len(arr); i++ {
		list.Insert(arr[i])
	}
	list.display()
	fmt.Println()
	list.head = list.Reverse(list.head, 3)
	list.display()
	fmt.Println()
}
