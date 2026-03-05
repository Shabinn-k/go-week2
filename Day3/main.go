package main
import f "fmt"
type Node struct{
	Data int
	next *Node
}
func main() {
	node1:=&Node{Data: 10}
	node2:=&Node{Data: 20}
	node3:=&Node{Data: 30}
	
	node1.next=node2
	node2.next=node3

	head:=node1
	current:=head

	for current!=nil{
		f.Println(current.Data)
		current=current.next
	}

}