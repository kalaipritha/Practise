package main

import "fmt"

func (p *node) RemoveDuplicate() {
	hash := make(map[int]int)
	cur := p.head
	nex := cur.next
	hash[cur.data] = 1
	for cur.next != nil {
		if hash[nex.data] == 0 {
			hash[nex.data] = 1
			cur = nex
			nex = nex.next
		} else {
			cur.next = nex.next
			nex = nex.next
		}
	}
}

func main() {
	list := createList()
	arr := []int{10, 12, 11, 11, 12, 11, 10}
	for i := 0; i < len(arr); i++ {
		list.Insert(arr[i])
	}
	list.display()
	fmt.Println()
	list.RemoveDuplicate()
	list.display()
	fmt.Println()
}
