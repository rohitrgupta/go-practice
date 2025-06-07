package main

import "fmt"

type ListNode struct {
	Value int
	Next  *ListNode
}

type LinkedList struct {
	Head *ListNode
}

func (l *LinkedList) Display() {
	if l.Head == nil {
		return
	}
	current := l.Head
	for current != nil {
		fmt.Print(current.Value, " ")
		current = current.Next
	}
	fmt.Println()
}

func (l *LinkedList) Length() int {
	length := 0
	current := l.Head
	for current != nil {
		length++
		current = current.Next
	}
	return length
}

func (l *LinkedList) InsertAtBeginning(value int) {
	newNode := &ListNode{Value: value}
	newNode.Next = l.Head
	l.Head = newNode
}

func (l *LinkedList) InsertAtEnd(value int) {
	newNode := &ListNode{Value: value}
	if l.Head == nil {
		l.Head = newNode
		return
	}
	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

func (l *LinkedList) InsertAfter(value, after int) {
	if l.Head == nil {
		return
	}
	current := l.Head
	for current != nil && current.Value != after {
		current = current.Next
	}
	if current != nil {
		newNode := &ListNode{Value: value}
		newNode.Next = current.Next
		current.Next = newNode
	}
}

func (l *LinkedList) Delete(value int) {
	if l.Head == nil {
		return
	}
	if l.Head.Value == value {
		l.Head = l.Head.Next
		return
	}
	current := l.Head
	for current.Next != nil && current.Next.Value != value {
		current = current.Next
	}
	if current.Next != nil {
		current.Next = current.Next.Next
	}
}

func (l *LinkedList) Search(value int) *ListNode {
	current := l.Head
	for current != nil {
		if current.Value == value {
			return current
		}
		current = current.Next
	}
	return nil
}

func (l *LinkedList) Reverse() {
	var prev *ListNode
	current := l.Head
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	l.Head = prev
}
