// package main

// import "fmt"

// type que struct {
// 	arr []int
// }

// func (q *que) Enque(val int) {
// 	q.arr = append(q.arr, val)
// }
// func (q *que) Deque() int {
// 	if len(q.arr) == 0 {
// 		fmt.Println("queue underflow")
// 		return -1
// 	}
//     val:=q.arr[0]
//     q.arr=q.arr[1:]
//     return val
// }
// func (q que)peek()int{
//     if len(q.arr)==0{
//         fmt.Println("underflow")
//         return -1
//     }
//     return q.arr[0]
// }
// func (q *que)display(){
//     fmt.Println(q.arr)
// }
// func main() {
// a:=que{}
// a.Enque(10)
// a.Enque(20)
// a.Enque(30)
// a.display()
// fmt.Println(a.Deque())
// fmt.Println(a.Deque())
// fmt.Println(a.Deque())
// fmt.Println(a.Deque())
// fmt.Println(a.peek())
// }

package main

import "fmt"
type Node struct{
    data int
    next *Node
}
type que struct{
    fist *Node
    last *Node
}
func (q *que)enq(val int){
    newNode:=&Node{data: val}
    if q.last==nil{
        q.fist=newNode
        q.last=newNode
        return
    }
    q.last.next=newNode
    q.last=newNode
}
func (q *que)deq()int{
    if q.fist==nil{
        fmt.Println("empty q")
        return -1
    }
    val:=q.fist.data
    q.fist=q.fist.next
    if q.fist==nil{
        q.last=nil
    }
    return val
}

func (q *que)display(){
    temp:=q.fist
    for temp!=nil{
        fmt.Println(temp.data)
        temp=temp.next
    }
}
func main() {
    s:=que{}
    s.enq(10)
    s.enq(20)
    s.enq(30)
    s.display()
    fmt.Println(s.deq())
    s.display()
}