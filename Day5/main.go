// package main

// import "fmt"

// type stack struct {
// 	arr []int
// }

// func (s *stack) push(val int) {
// 	s.arr = append(s.arr, val)
// }
// func (s *stack) pop() int {
// 	if len(s.arr) == 0 {
// 		fmt.Println("stack underflow")
// 		return -1
// 	}
// 	top:=s.arr[len(s.arr)-1]
// 	s.arr=s.arr[:len(s.arr)-1]
// 	return top
// }
// func (s *stack)peek()int{
// return s.arr[len(s.arr)-1]
// }
// func (s *stack) display() {
// 	fmt.Println(s.arr)
// }
// func main() {
// a:=stack{}
// a.push(10)
// a.push(20)
// a.push(30)
// a.push(40)
// a.display()
// fmt.Println(a.pop())
// fmt.Println(a.peek())
// }

package main

import "fmt"
type Node struct{
	data int
	next *Node
}
type Stack struct{
	top *Node
}
func (s *Stack)push(val int){
	newNode:=&Node{data :val}
	newNode.next=s.top
	s.top=newNode
}
func (s *Stack)pop()int{
	if s.top==nil{
		fmt.Println("stack underflow")
		return -1
	}
	val:=s.top.data
	s.top=s.top.next
	return val
}
func (s *Stack)peek()int{
	return s.top.data
}
func (s *Stack)display(){
	temp:=s.top
	for temp!=nil{
		fmt.Println(temp.data)
		temp=temp.next
	}
}
func main() {
	a:=Stack{}
	a.push(10)
	a.push(20)
	a.push(30)
	a.display()
	fmt.Println(a.pop())
	fmt.Println(a.peek())
}