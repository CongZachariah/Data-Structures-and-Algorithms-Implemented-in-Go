package main

import "fmt"

func main() {
	fmt.Print(reversePairs([]int {5,3,1,7,5,3,2,4,2}))
}

func reversePairs(nums []int) int {
	res := 0
	for i := 0 ;i < len(nums) ; i ++ {
		for j := i ; j <len(nums) ;j ++ {
			if nums[i] > nums[j] {
				res ++
			}
		}
	}
	return res
}