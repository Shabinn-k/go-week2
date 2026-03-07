// package main
// import f "fmt"
// func change(n *int){
// 	*n=100
// }

// func main() {
// 	x:=10
// 	change(&x)
// 	f.Println(x)

// }

package main

import "fmt"

func main() {
	// x:=10
	// p:=&x
	// pp:=&p
	// fmt.Println(**pp)

	x:=20
	p:=&x
	fmt.Println(*p)
	*p=40
	fmt.Println(*p)

}