package main

import "fmt"

type list struct {
	data int
	next *list
}

type node struct {
	head        *list
	currentNode *list
}

func createList() *node {
	return &node{}
}

func (p *node) Insert(data int) error {
	s := &list{
		data: data,
	}
	if p.head == nil {
		p.head = s
		p.currentNode = s
	} else {
		currentNode := p.currentNode
		currentNode.next = s
		p.currentNode = s
	}
	return nil
}

func (p *node) display() error {
	currentNode := p.head
	if currentNode == nil {
		fmt.Println("list is empty.")
		return nil
	}
	fmt.Print(currentNode.data, " -> ")
	for currentNode.next != nil {
		currentNode = currentNode.next
		fmt.Print(currentNode.data, " -> ")
	}

	return nil
}

func (p *node) delete(n int) {
	curr := p.head

	if curr.data == n {
		p := curr.next
		curr.next = p.next
		p.next = nil
	}
}
