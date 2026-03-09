// package main
// import "fmt"
// func main() {
// 	arr := []int{1, 2, 3, 4}
// 	arr = append(arr, 5)
// 	arr = append(arr, 0)
// 	val:=6
// 	index:=5
// 	for i:=len(arr)-1;i>index;i--{
// 		arr[i]=arr[i-1]
// 	}
// 	arr[index]=val
// 	arr = append(arr[:1],arr[2:]...)
// 	fmt.Println(arr)

// 	target:=4
// 	for i,v :=range arr{
// 		if v==target{
// 			fmt.Println(i)
// 		}
// 	}
// 	max:=arr[0]
// 	min:=arr[0]
// 	for i:=0;i<len(arr)-1;i++{
// 		if arr[i]>max{
// 			max=arr[i]
// 		}else{
// 			min=arr[i]
// 		}
// 	}
// 	fmt.Println(max)
// 	fmt.Println(min)
// 	f:=0
// 	l:=len(arr)-1
// 	for f<l{
// 		arr[f],arr[l]=arr[l],arr[f]
// 		f++
// 		l--
// 	}
// 	fmt.Println(arr)
// }

// package main
// import "fmt"
// func main() {
// 	arr:=[]int{1,2,3,4,5}
// 	// f:=arr[0]
// 	// for i:=0;i<len(arr)-1;i++{
// 	// 	arr[i]=arr[i+1]
// 	// }
// 	// arr[len(arr)-1]=f
// 	// fmt.Println(arr)

// 	f:=arr[len(arr)-1]
// 	for i:=len(arr)-1;i>0;i--{
// 		arr[i]=arr[i-1]
// 	}
// 	arr[0]=f
// 	fmt.Println(arr)
// }

// package main
// import (
// 	"sort"
// 	"fmt"
// )
// func main() {
	// arr:=[]int{1,1,2,3,3,4,5,6,6}
	// m:=make(map[int]int)
	// dupe:=[]int{}
	// for _,v:=range arr{
	// 	m[v]++
	// if m[v]>1{
	// 	dupe = append(dupe, v)
	// }
	// }
	// fmt.Println(dupe)
	// fmt.Println(m)

	// arr:=[]int{0,1,0,2,3,4,0,5,6,0,7}
	// j:=0
	// for i:=range arr{
	// 	if arr[i]!=0{
	// 		arr[i],arr[j]=arr[j],arr[i]
	// 		j++
	// 	}
	// }
	// fmt.Println(arr)

	// arr:=[]int{3,2,4,1,5,45,2,52,6,6}
	// sort.Ints(arr)
	// fmt.Println(arr)
	// for i:=1;i<len(arr);i++{
	// 	if arr[i]<arr[i-1]{
	// 		fmt.Println("it not is sorted")
	// 		return 
	// 	}
	// }
	// fmt.Println("it is sorted")
// }

// package main
// import "fmt"
// func main() {
// 	arr:=[]int{2,7,11,15}
// 	target:=9
// 	m:=make(map[int]int)
// 	for i:=0;i<len(arr);i++{
// 		rem:=target-arr[i]
// 		if _,ok:=m[rem];ok{
// 			fmt.Println(rem,arr[i])
// 			return
// 		}
// 		m[arr[i]]=i
// 	}
// }

//single linked
package main
import "fmt"
type Node struct{
	data int
	next *Node
}
func main() {
	node1:=&Node{data: 10}
	node2:=&Node{data: 20}
	node3:=&Node{data: 30}

	node1.next=node2	
	node2.next=node3

	head:=node1
	current:=head

	node0:=&Node{data: 0}
	node0.next=head
	head=node0

	newNode:=&Node{data: 40}
	for current.next!=nil{
		current=current.next
	}
	current.next=newNode
	
	for current.next.next!=nil{
		current=current.next
	}
	current.next=nil
	head=head.next
	temp:=head
	for temp!=nil{
		fmt.Println(temp.data)
		temp=temp.next
	}
}

//double linked
// package main
// import "fmt"
// type Node struct{
// 	prev *Node
// 	data int
// 	next *Node
// }
// func main() {
// 	node1:=&Node{data: 10}
// 	node2:=&Node{data: 20}
// 	node3:=&Node{data: 30}

// 	node1.next=node2
// 	node2.prev=node1	
// 	node2.next=node3
// 	node3.prev=node2

// 	head:=node1
// 	current:=head

// 	node0:=&Node{data: 0}
// 	node0.next=head
// 	head.prev=node0
// 	head=node0
// 	newNode:=&Node{data: 40}
// 	for current.next!=nil{
// 		current=current.next
// 	}
// 	current.next=newNode
// 	newNode.prev=current
// 	current.next=nil
// 	head=head.next
// 	head.prev=nil
// 	for head!=nil{
// 		fmt.Println(head.data)
// 		head=head.next
// 	}
// }