package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	for i := 0; i < len(nums); i++ {
		for k := i + 1; k < len(nums); k++ {
			if nums[i]+nums[k] == target {
				fmt.Println([]int{i,k})
			}
		}
	}
	fmt.Println(nil)
}

