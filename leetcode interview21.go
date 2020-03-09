package main

import "fmt"

func main() {
	fmt.Print(exchange([]int {1,7,4,3,5,3,5,3,2,7,8,7,5}))
}

func exchange(nums []int) []int {
	numjishu := []int{}
	numoushu := []int{}

	for i:=0; i<len(nums); i++ {
		if nums[i]%2 == 1 {
			numjishu = append(numjishu, nums[i])
		}else{
			numoushu = append(numoushu, nums[i])
		}
	}
	return append(numjishu, numoushu...)
}