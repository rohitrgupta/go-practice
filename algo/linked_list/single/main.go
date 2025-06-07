package main

import "fmt"

func main() {
	// Example usage of the linked list implementation
	list := LinkedList{
		Head: &ListNode{Value: 1},
	}
	// list.Head.Next = &ListNode{Value: 2}
	list.InsertAtBeginning(2)
	list.InsertAtBeginning(3)
	list.InsertAtEnd(4)
	fmt.Println("Length of the linked list:", list.Length())
	list.Display()
	list.InsertAfter(5, 2)
	list.Display()
	list.Delete(3)
	list.Display()
	fmt.Printf("searching for 4: %v\n", list.Search(4))
}
