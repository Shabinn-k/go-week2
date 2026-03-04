// package main
// import f "fmt"
// func change(n *int){
// 	*n=100
// }
// func main() {
// 	x:=10
// 	f.Println("before changing value -",x)
// 	change(&x)
// 	f.Println("After changing value -",x)
// }

// package main
// import "fmt"
// func main() {

// 	// Using new
// 	ptr := new(int)
// 	fmt.Println("Value from new:", *ptr)

// 	// Using make for slice
// 	s := make([]int, 3)
// 	fmt.Println("Slice:", s)

// 	// Using make for map
// 	m := make(map[string]int)
// 	m["Shabin"] = 1
// 	fmt.Println("Map:", m)

// 	// What happens if we use new for map?
// 	mp := new(map[string]int)
// 	// (*mp)["fail"] = 1  // This will panic
// 	fmt.Println("Map using new:", mp)
// }

package main

import f "fmt"
func Maxarr(num []int)int{
	max:=num[0]
	for i:=range num{
		if num[i]>max{
			max=num[i]
		}
	}
	return  max
}
func main() {
	arr:=[]int{10,2,41,5,523,654,21,9}
	f.Println("max number in array is",Maxarr(arr))
}