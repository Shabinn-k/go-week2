//INSERT AT END
// package main
// import f "fmt"
// type Node struct{
// 	data int
// 	next *Node
// }
// func main() {
// node1:=&Node{data: 1}
// node2:=&Node{data: 2}
// node3:=&Node{data: 3}

// node1.next=node2
// node2.next=node3

// head:=node1
// current:=head

// 	node4:=&Node{data: 4}
// 	for current.next!=nil{
// 		current=current.next
// 	}
// 	current.next=node4

// 	node5:=&Node{data: 5}
// 	for current.next!=nil{
// 		current=current.next
// 	}
// 	current.next=node5

// 	for head!=nil{
// 		f.Println(head.data)
// 		head=head.next
// 	}

// }

//INSERT AT BEGINNING
// package main
// import f "fmt"
// type Node struct{
// 	data int
// 	next *Node
// }
// func main() {
// 	node2:=&Node{data: 2}
// 	node3:=&Node{data: 3}
// 	node4:=&Node{data: 4}

// 	node2.next=node3
// 	node3.next=node4

// 	head:=node2

// 	node1:=&Node{data: 1}
// 	node1.next=head
// 	head=node1

// 	node0:=&Node{data: 0}
// 	node0.next=head
// 	head=node0

// 	current:=head

// 	for current!=nil{
// 		f.Println(current.data)
// 		current=current.next
// 	}
// }

//DELETING AT END
// package main
// import f "fmt"
// type Node struct{
// 	data int
// 	next *Node
// }
// func main() {
// 	node1:=&Node{data: 1}
// 	node2:=&Node{data: 2}
// 	node3:=&Node{data: 3}

// 	node1.next=node2
// 	node2.next=node3

// 	head:=node1
// 	current:=head

// 	for current.next.next!=nil{
// 		current=current.next
// 	}
// 	current.next=nil

// 	for head!=nil{
// 		f.Println(head.data)
// 		head=head.next
// 	}
// }

//DELETING AT BEGINNING
// package main
// import f "fmt"
// type Node struct{
// 	data int
// 	next *Node
// }
// func main() {
// 	node1:=&Node{data: 1}
// 	node2:=&Node{data: 2}
// 	node3:=&Node{data: 3}

// 	node1.next=node2
// 	node2.next=node3

// 	head:=node1
// 	head=head.next

// 	for head!=nil{
// 		f.Println(head.data)
// 		head=head.next
// 	}
// }

// package main
// import "fmt"
// type Node struct {
// 	data int
// 	next *Node
// }

// func main() {
// 	node1 := &Node{data: 10}
// 	node2 := &Node{data: 20}
// 	node3 := &Node{data: 30}

// 	node1.next = node2
// 	node2.next = node3

// 	head := node1
// 	current := head
// 	target:=20

// 	for current != nil {
// 		if current.data==target{
// 			fmt.Println("value found",current.data)
// 			return
// 		}
// 		current=current.next
// 	}
// 	fmt.Println("value not found")

// }

// package main
// import "fmt"
// type Node struct {
// 	data int
// 	next *Node
// }

// func main() {
// 	node1 := &Node{data: 10}
// 	node2 := &Node{data: 20}
// 	node3 := &Node{data: 30}

// 	node1.next = node2
// 	node2.next = node3

// 	head := node1
// 	current:=head

// 	 var arr []int
// 	 for current!=nil{
// 		arr = append(arr, current.data)
// 		current=current.next
// 	 }
// 	 fmt.Println(arr)
// }

package main

import "fmt"
type Node struct {
	data int
	next *Node
}
func main() {
	arr:=[]int{1,2,3,4,5}

	head:=&Node{data: arr[0]}
	current:=head

	for i:=1;i<len(arr);i++{
		newNode:=&Node{data: arr[i]}
		current.next=newNode
		current=newNode
	}
	temp:=head
	for temp!=nil{
		fmt.Println(temp.data)
		temp=temp.next
	}
}