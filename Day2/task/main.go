// 1. Insert/Delete at specific index

// package main
// import f "fmt"
// func main() {
// 	array:=[]int{10,20,30,50,60,70}
// 	array = append(array,0)

// 	index:=3
// 	value:=40
// 	for i:=len(array)-1;i>index;i--{
// 		array[i]=array[i-1]
// 	}
// 	array[index]=value
// 	f.Println(array)	

// 	array = append(array[:4],array[5:]...)
// 	f.Println(array)
// }


// 2. Reverse an array
// 3. Solve 2–3 LeetCode problems (Two Sum, Merge Sorted Arrays, etc.)
package main
import "fmt"
func main() {
	array:=[]int{5,4,3,2,1}
	f:=0
	l:=len(array)-1
	for l>f{
		array[f],array[l]=array[l],array[f]
		f++
		l--
	}
	fmt.Println(array)
}