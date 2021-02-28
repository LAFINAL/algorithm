package main

import "fmt"

type linkedlist struct {
	head *node
	tail *node
}

type node struct {
	value int
	next *node
}

func main() {
	a := new(linkedlist)
	a.push(5)
	a.push(19)
	fmt.Println(a.first().value)
	fmt.Println(a.first().next.value)

}

func (l *linkedlist) push(value int) {
	node := &node{value: value}
	if l.head == nil {
		l.head = node
	} else {
		l.tail.next = node
	}
	l.tail = node
}

func (l *linkedlist) first() *node{
	return l.head
}

func (l *node) nextValue() *node {
	return l.next
}