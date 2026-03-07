package main

import "fmt"

type Node struct {
	data int
	next *Node
}

func main() {
	node2 := &Node{data: 20}
	node3 := &Node{data: 30}
	node4 := &Node{data: 40}

	node2.next = node3
	node3.next = node4

	head := node2

	node1 := &Node{data: 10}
	node1.next = head
	head = node1

	current := head
	node5 := &Node{data: 50}
	for current.next != nil {
		current = current.next
	}
	current.next = node5

	for current.next.next != nil {
		current = current.next
	}
	current.next = nil

	head = head.next
	for head != nil {
		fmt.Println(head.data)
		head = head.next
	}

}
