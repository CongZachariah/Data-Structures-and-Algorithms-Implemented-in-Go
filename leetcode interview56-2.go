package main

import (
	"fmt"
	"sort"
)

func main() {
	data := []int {2,2,4,3,1,2,6,4,4,3,8,3,1,1,6,6}
	fmt.Print(singleNumber(data))
}

func singleNumber(nums []int) int {
	sort.Ints(nums)
	for i:=0; i<len(nums)-3 ; i += 3{
		if nums[i] == nums[i+1] && nums[i] == nums[i+2]{
			continue
		} else {
			return nums[i]
		}
	}
    return nums[ len(nums) - 1 ]
}