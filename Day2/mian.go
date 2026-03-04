// package main
// import f "fmt"
// func main() {
// 	arr:=[]int{1,3,4,5,6}
// 	f.Println("capacity at first",cap(arr))

// 	// arr = append(arr, 7)
// 	// f.Println(arr)
// 	// f.Println(cap(arr))
	
// 	arr=append(arr, 0)
// 	f.Println(arr)
// 	f.Println("capcaity at second",cap(arr))
	
// 	ind:=1
// 	val:=2
// 	for i:=len(arr)-1;i>0;i--{
// 		arr[i]=arr[i-1]
// 	}
// 	arr[ind]=val

// 	f.Println(arr)
// 	f.Println("capacity at last",cap(arr))

// 	// arr=append(arr[:1],arr[2:]...)
// 	// f.Println(arr)
	
// }

package main
import f "fmt"
func main() {
	arr:=[]int{10,20,30,40,50}
	tg:=20

	//Linear search
	for k,v:=range arr{
		if v==tg{
			f.Println(k)
		}
	}

	//Binary search
	l:=0
	r:=len(arr)-1
	for l<=r{
		mid:=(l+r)/2
		if arr[mid]==tg{
			f.Println(mid)
		}
		if arr[mid]<tg{
		l=mid+1	
		}else{
			r=mid-1
		}
	}
}